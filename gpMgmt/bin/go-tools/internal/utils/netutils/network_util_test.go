package netutils_test

import (
	"testing"

	"github.com/greenplum-db/gpdb/gp/internal/utils/netutils"
	"github.com/greenplum-db/gpdb/gp/testutils"
)

func TestNetworkUtil(t *testing.T) {

	t.Run("IsValidIpv4", func(t *testing.T) {

		netUtil := netutils.NewNetworkUtil()

		isValidIpv4 := netUtil.IsValidIpv4("192.168.12.7")
		testutils.Assert(t, true, isValidIpv4, "")

		isValidIpv4 = netUtil.IsValidIpv4("192.168.12.7.32")
		testutils.Assert(t, false, isValidIpv4, "")

		isValidIpv4 = netUtil.IsValidIpv4("192:168.12.7")
		testutils.Assert(t, false, isValidIpv4, "")

		isValidIpv4 = netUtil.IsValidIpv4("192.168.12.7:80")
		testutils.Assert(t, false, isValidIpv4, "")

		isValidIpv4 = netUtil.IsValidIpv4("2001:db8::1234:5678")
		testutils.Assert(t, false, isValidIpv4, "")
	})

	t.Run("generate ipv4 range", func(t *testing.T) {

		netUtil := netutils.NewNetworkUtil()
		ipList, err := netUtil.GenerateIpv4List("192.168.1.7", "192.168.1.11")

		testutils.Assert(t, nil, err, "")

		expectedIpList := []string{
			"192.168.1.7",
			"192.168.1.8",
			"192.168.1.9",
			"192.168.1.10",
			"192.168.1.11",
		}

		testutils.Assert(t, expectedIpList, ipList, "")

	})

}
