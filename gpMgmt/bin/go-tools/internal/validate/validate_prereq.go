package validate

import (
	"fmt"

	"github.com/greenplum-db/gpdb/gp/internal/enums"
)

type prereqStageValidator struct{}

func NewPrereqStageValidator() StageValidator {
	return &prereqStageValidator{}
}

func (prereqStageValidator) GetValidators() []Validator {
	return []Validator{
		NewCoordinatorHostValidator(),
		NewNetworkValidator(),
		NewArtifactValidator(),
	}
}

func (prereqStageValidator) ValidateConfig() error {
	conf, err := Dependency.GetConfig()
	if err != nil {
		return err
	}

	netUtil := Dependency.NewNetworkUtil()

	// coordinator config
	switch true {
	case conf.GetInfraConfig() == nil:
		return fmt.Errorf("infra config not found")
	case conf.GetInfraConfig().GetCoordinator() == nil:
		return fmt.Errorf("coordinator config not found")
	case !netUtil.IsValidIpv4(conf.GetInfraConfig().GetCoordinator().GetIp()):
		return fmt.Errorf("coordinator ip not valid IPv4")
	}

	// standby config
	switch true {
	case conf.GetDatabaseConfig() == nil:
		return fmt.Errorf("database config not found")
	case conf.GetDatabaseConfig().GetDeploymentType() == enums.DeploymentTypeMirrored && conf.GetInfraConfig().GetStandby() == nil:
		return fmt.Errorf("standby config not found")
	case conf.GetDatabaseConfig().GetDeploymentType() == enums.DeploymentTypeMirrored && !netUtil.IsValidIpv4(conf.GetInfraConfig().GetStandby().GetIp()):
		return fmt.Errorf("standby ip not valid IPv4")
	}

	// segmenthost config
	switch true {
	case conf.GetInfraConfig().GetSegmentHost() == nil:
		return fmt.Errorf("segmenthost config not found")
	case conf.GetInfraConfig().GetSegmentHost().GetNetwork() == nil:
		return fmt.Errorf("segmenthost network config not found")
	}

	// segmenthost ip list
	if err := conf.GenerateSegmentIPList(); err != nil {
		return err
	}
	if len(conf.GetInfraConfig().GetSegmentHost().GetNetwork().GetIpList()) == 0 {
		return fmt.Errorf("segmenthost ip range/list config not found")
	}

	// artifact config
	if conf.GetArtifactConfig() == nil {
		return fmt.Errorf("artifact config not found")
	}

	return nil
}
