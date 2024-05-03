package config

import (
	"fmt"

	"github.com/greenplum-db/gpdb/gp/internal/utils/netutils"
	"github.com/spf13/viper"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Config
type Config interface {
	SetName(string)
	Load(string) error

	GetDatabaseConfig() DatabaseConfig
	GetArtifactConfig() ArtifactConfig
	GetInfraConfig() InfraConfig

	GenerateSegmentIPList() error
}

func New() Config {
	return &appConfig{
		configName: "gp.conf",
	}
}

type appConfig struct {
	configName string
	// TODO : embed hub Config

	Database *Database `json:"database"`
	Artifact *Artifact `json:"artifact"`
	Infra    *Infra    `json:"infra"`
}

func (conf *appConfig) SetName(configName string) {
	conf.configName = configName
}

func (conf *appConfig) Load(configPath string) error {
	parser := viper.New()
	parser.SetConfigName(conf.configName)
	parser.SetConfigType("json")

	parser.AddConfigPath(configPath)

	conf.setDefaults(parser)
	err := parser.ReadInConfig()
	if err != nil {
		return err
	}

	err = parser.Unmarshal(conf)
	if err != nil {
		return err
	}

	return nil
}

func (conf *appConfig) GetDatabaseConfig() DatabaseConfig {
	return conf.Database
}

func (conf *appConfig) GetArtifactConfig() ArtifactConfig {
	return conf.Artifact
}

func (conf *appConfig) GetInfraConfig() InfraConfig {
	return conf.Infra
}

func (conf *appConfig) setDefaults(parser *viper.Viper) *viper.Viper {
	parser.SetDefault("Infra.RequestPort", 4506)
	parser.SetDefault("Infra.PublishPort", 4505)
	parser.SetDefault("Infra.Coordinator.HostName", "cdw")
	parser.SetDefault("Database.Admin.Name", "gpadmin")

	return parser
}

func (conf *appConfig) GenerateSegmentIPList() error {
	netUtil := netutils.NewNetworkUtil()
	if conf.Infra == nil || conf.Infra.SegmentHosts == nil || conf.Infra.SegmentHosts.Network == nil {
		return fmt.Errorf("segment network config not found")
	}
	if len(conf.Infra.SegmentHosts.Network.IpList) > 0 {
		return nil
	}
	if conf.Infra.SegmentHosts.Network.IpRange == nil {
		return fmt.Errorf("segment ip range/list not found")
	}
	firstIp := conf.GetInfraConfig().GetSegmentHost().GetNetwork().GetIpRange().GetFirstIp()
	lastIp := conf.GetInfraConfig().GetSegmentHost().GetNetwork().GetIpRange().GetLastIp()
	ipList, err := netUtil.GenerateIpv4List(firstIp, lastIp)
	if err != nil {
		return err
	}
	conf.Infra.SegmentHosts.Network.IpList = ipList
	return nil
}
