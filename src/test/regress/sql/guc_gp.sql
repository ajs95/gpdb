SELECT min_val, max_val FROM pg_settings WHERE name = 'gp_resqueue_priority_cpucores_per_segment';

-- Test cursor gang should not be reused if SET command happens.
CREATE OR REPLACE FUNCTION test_set_cursor_func() RETURNS text as $$
DECLARE
  result text;
BEGIN
  EXECUTE 'select setting from pg_settings where name=''temp_buffers''' INTO result;
  RETURN result;
END;
$$ LANGUAGE plpgsql;

SET temp_buffers = 2000;
BEGIN;
  DECLARE set_cusor CURSOR FOR SELECT relname FROM gp_dist_random('pg_class');
  -- The GUC setting should not be dispatched to the cursor gang.
  SET temp_buffers = 3000;
END;

-- Verify the cursor gang is not reused. If the gang is reused, the
-- temp_buffers value on that gang should be old one, i.e. 2000 instead of
-- the new committed 3000.
SELECT * from (SELECT test_set_cursor_func() FROM gp_dist_random('pg_class') limit 1) t1
  JOIN (SELECT test_set_cursor_func() FROM gp_dist_random('pg_class') limit 1) t2 ON TRUE;

RESET temp_buffers;

--
-- Test GUC if cursor is opened
--
-- start_ignore
drop table if exists test_cursor_set_table;
drop function if exists test_set_in_loop();
drop function if exists test_call_set_command();
-- end_ignore

create table test_cursor_set_table as select * from generate_series(1, 100);

CREATE FUNCTION test_set_in_loop () RETURNS numeric
    AS $$
DECLARE
    rec record;
    result numeric;
    tmp numeric;
BEGIN
	result = 0;
FOR rec IN select * from test_cursor_set_table
LOOP
        select test_call_set_command() into tmp;
        result = result + 1;
END LOOP;
return result;
END;
$$
    LANGUAGE plpgsql NO SQL;


CREATE FUNCTION test_call_set_command() returns numeric
AS $$
BEGIN
       execute 'SET gp_workfile_limit_per_query=524;';
       return 0;
END;
$$
    LANGUAGE plpgsql NO SQL;

SELECT * from test_set_in_loop();


CREATE FUNCTION test_set_within_initplan () RETURNS numeric
AS $$
DECLARE
	result numeric;
	tmp RECORD;
BEGIN
	result = 1;
	execute 'SET gp_workfile_limit_per_query=524;';
	select into tmp * from test_cursor_set_table limit 100;
	return result;
END;
$$
	LANGUAGE plpgsql;


CREATE TABLE test_initplan_set_table as select * from test_set_within_initplan();


DROP TABLE if exists test_initplan_set_table;
DROP TABLE if exists test_cursor_set_table;
DROP FUNCTION if exists test_set_in_loop();
DROP FUNCTION if exists test_call_set_command();


-- Set work_mem. It emits a WARNING, but it should only emit it once.
--
-- We used to erroneously set the GUC twice in the QD node, whenever you issue
-- a SET command. If this stops emitting a WARNING in the future, we'll need
-- another way to detect that the GUC's assign-hook is called only once.
set work_mem='1MB';
reset work_mem;

--
-- Test if RESET timezone is dispatched to all slices
--
CREATE TABLE timezone_table AS SELECT * FROM (VALUES (123,1513123564),(123,1512140765),(123,1512173164),(123,1512396441)) foo(a, b) DISTRIBUTED RANDOMLY;

SELECT to_timestamp(b)::timestamp WITH TIME ZONE AS b_ts FROM timezone_table ORDER BY b_ts;
SET timezone= 'America/New_York';
-- Check if it is set correctly on QD.
SELECT to_timestamp(1613123565)::timestamp WITH TIME ZONE;
-- Check if it is set correctly on the QEs.
SELECT to_timestamp(b)::timestamp WITH TIME ZONE AS b_ts FROM timezone_table ORDER BY b_ts;
RESET timezone;
-- Check if it is reset correctly on QD.
SELECT to_timestamp(1613123565)::timestamp WITH TIME ZONE;
-- Check if it is reset correctly on the QEs.
SELECT to_timestamp(b)::timestamp WITH TIME ZONE AS b_ts FROM timezone_table ORDER BY b_ts;

--
-- Test if SET TIME ZONE INTERVAL is dispatched correctly to all segments
--
SET TIME ZONE INTERVAL '04:30:06' HOUR TO MINUTE;
-- Check if it is set correctly on QD.
SELECT to_timestamp(1613123565)::timestamp WITH TIME ZONE;
-- Check if it is set correctly on the QEs.
SELECT to_timestamp(b)::timestamp WITH TIME ZONE AS b_ts FROM timezone_table ORDER BY b_ts;

-- Test default_transaction_isolation and transaction_isolation fallback from serializable to repeatable read
CREATE TABLE test_serializable(a int);
insert into test_serializable values(1);
SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL serializable;
show default_transaction_isolation;
SELECT * FROM test_serializable;
SET default_transaction_isolation = 'read committed';
SET default_transaction_isolation = 'serializable';
show default_transaction_isolation;
SELECT * FROM test_serializable;
SET default_transaction_isolation = 'read committed';

BEGIN TRANSACTION ISOLATION LEVEL serializable;
	show transaction_isolation;
	SELECT * FROM test_serializable;
COMMIT;
DROP TABLE test_serializable;
--
-- Test DISCARD TEMP.
--
-- There's a test like this in upstream 'guc' test, but this expanded version
-- verifies that temp tables are dropped on segments, too.
--
CREATE TEMP TABLE reset_test ( data text ) ON COMMIT DELETE ROWS;
DISCARD TEMP;
-- Try to create a new temp table with same. Should work.
CREATE TEMP TABLE reset_test ( data text ) ON COMMIT PRESERVE ROWS;

-- Now test that the effects of DISCARD TEMP can be rolled back
BEGIN;
DISCARD TEMP;
ROLLBACK;
-- the table should still exist.
INSERT INTO reset_test VALUES (1);

-- Unlike DISCARD TEMP, DISCARD ALL cannot be run in a transaction.
BEGIN;
DISCARD ALL;
COMMIT;
-- the table should still exist.
INSERT INTO reset_test VALUES (2);
SELECT * FROM reset_test;

DISCARD ALL;
CREATE TEMP TABLE reset_test ( data text ) ON COMMIT PRESERVE ROWS;
