package validate

import (
	"fmt"
	"sync"

	"github.com/greenplum-db/gpdb/gp/internal/enums"
)

type coordinatorHostValidator struct{}

func NewCoordinatorHostValidator() Validator {
	return &coordinatorHostValidator{}
}

func (v coordinatorHostValidator) Validate() error {
	conf, err := Dependency.GetConfig()
	if err != nil {
		return err
	}

	hostUtil := Dependency.NewHostUtil()
	if err := hostUtil.VerifyHostIp(conf.GetInfraConfig().GetCoordinator().GetIp()); err != nil {
		return err
	}

	if conf.GetDatabaseConfig().GetDeploymentType() == enums.DeploymentTypeMirrored {
		if err := hostUtil.VerifyHostIp(conf.GetInfraConfig().GetStandby().GetIp()); err != nil {
			return err
		}
	}
	return nil
}

type networkValidator struct{}

func NewNetworkValidator() Validator {
	return &networkValidator{}
}

func (v networkValidator) Validate() error {
	conf, err := Dependency.GetConfig()
	if err != nil {
		return err
	}

	netUtil := Dependency.NewNetworkUtil()
	var wg sync.WaitGroup
	var hostsNotReachable []string
	var writeLock sync.Mutex
	for _, segmentIp := range conf.GetInfraConfig().GetSegmentHost().GetNetwork().GetIpList() {
		wg.Add(1)
		go func(segIp string) {
			defer wg.Done()
			if !netUtil.IsReachable(segIp) {
				writeLock.Lock()
				hostsNotReachable = append(hostsNotReachable, segIp)
				writeLock.Unlock()
			}
		}(segmentIp)
	}
	wg.Wait()

	if len(hostsNotReachable) > 0 {
		return fmt.Errorf("segment hosts not reachable : %v", hostsNotReachable)
	}
	return nil
}
