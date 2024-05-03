package utils

import "github.com/greenplum-db/gpdb/gp/internal/utils/netutils"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

var Dependency struct {
	NewNetworkUtil func() netutils.NetworkUtil
}

func init() {
	Dependency.NewNetworkUtil = netutils.NewNetworkUtil
}
