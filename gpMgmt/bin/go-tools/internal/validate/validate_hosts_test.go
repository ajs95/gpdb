package validate_test

import (
	"fmt"
	"testing"

	"github.com/greenplum-db/gpdb/gp/internal/config"
	"github.com/greenplum-db/gpdb/gp/internal/config/configfakes"
	"github.com/greenplum-db/gpdb/gp/internal/enums"
	"github.com/greenplum-db/gpdb/gp/internal/utils"
	"github.com/greenplum-db/gpdb/gp/internal/utils/utilsfakes"
	"github.com/greenplum-db/gpdb/gp/internal/validate"
	"github.com/greenplum-db/gpdb/gp/testutils"
)

func TestHostValidator(t *testing.T) {

	t.Run("Validate cordinator host fails when IP does not match", func(t *testing.T) {
		testHostUtil := &utilsfakes.FakeHostUtil{}
		validate.Dependency.NewHostUtil = func() utils.HostUtil {
			return testHostUtil
		}
		testHostUtil.VerifyHostIpReturns(fmt.Errorf("invalid host ip"))

		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}
		infraConfig := &configfakes.FakeInfraConfig{}
		testConfig.GetInfraConfigReturns(infraConfig)

		infraConfig.GetCoordinatorReturns(&configfakes.FakeHostConfig{})

		cdwValidator := validate.NewCoordinatorHostValidator()

		err := cdwValidator.Validate()

		testutils.AssertNotEmpty(t, err, "")
		testutils.Assert(t, "invalid host ip", err.Error(), "")
	})

	t.Run("Validate cordinator host succeeds", func(t *testing.T) {
		testHostUtil := &utilsfakes.FakeHostUtil{}
		validate.Dependency.NewHostUtil = func() utils.HostUtil {
			return testHostUtil
		}
		testHostUtil.VerifyHostIpReturns(nil)

		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}

		testDbConfig := &configfakes.FakeDatabaseConfig{}
		testConfig.GetDatabaseConfigReturns(testDbConfig)
		testDbConfig.GetDeploymentTypeReturns(enums.DeploymentTypeMirrored)

		infraConfig := &configfakes.FakeInfraConfig{}
		testConfig.GetInfraConfigReturns(infraConfig)
		infraConfig.GetCoordinatorReturns(&configfakes.FakeHostConfig{})
		infraConfig.GetStandbyReturns(&configfakes.FakeHostConfig{})

		cdwValidator := validate.NewCoordinatorHostValidator()

		err := cdwValidator.Validate()

		testutils.Assert(t, nil, err, "")
	})
}
