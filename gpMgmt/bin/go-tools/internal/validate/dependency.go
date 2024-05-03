package validate

import (
	"github.com/greenplum-db/gpdb/gp/internal/config"
	"github.com/greenplum-db/gpdb/gp/internal/utils"
	"github.com/greenplum-db/gpdb/gp/internal/utils/fsutils"
	"github.com/greenplum-db/gpdb/gp/internal/utils/netutils"
)

var Dependency struct {
	GetConfig         func() (config.Config, error)
	GetStageValidator func(Stage) (StageValidator, error)
	NewHostUtil       func() utils.HostUtil
	NewNetworkUtil    func() netutils.NetworkUtil
	NewFileSystem     func() fsutils.FileSystem
}

func init() {
	// TODO : Initialize GetConfig
	Dependency.GetStageValidator = getStageValidator
	Dependency.NewHostUtil = utils.NewHostUtil
	Dependency.NewNetworkUtil = netutils.NewNetworkUtil
	Dependency.NewFileSystem = fsutils.NewFileSystem
}
