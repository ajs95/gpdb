package validate_test

import (
	"fmt"
	"testing"

	"github.com/greenplum-db/gpdb/gp/internal/config"
	"github.com/greenplum-db/gpdb/gp/internal/config/configfakes"
	"github.com/greenplum-db/gpdb/gp/internal/enums"
	"github.com/greenplum-db/gpdb/gp/internal/validate"
	"github.com/greenplum-db/gpdb/gp/testutils"
)

func TestValidatePrereq(t *testing.T) {

	t.Run("Validate config fails when Infra config not present", func(t *testing.T) {

		testStageVldtr := validate.NewPrereqStageValidator()
		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}

		err := testStageVldtr.ValidateConfig()

		testutils.Assert(t, "infra config not found", err.Error(), "")
	})

	t.Run("Validate config fails when coordinator config not present", func(t *testing.T) {

		testStageVldtr := validate.NewPrereqStageValidator()
		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}
		testConfig.GetInfraConfigReturns(&configfakes.FakeInfraConfig{})

		err := testStageVldtr.ValidateConfig()

		testutils.Assert(t, "coordinator config not found", err.Error(), "")
	})

	t.Run("Validate config fails when coordinator ip is not valid", func(t *testing.T) {

		testStageVldtr := validate.NewPrereqStageValidator()
		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}
		infraConfig := &configfakes.FakeInfraConfig{}
		testConfig.GetInfraConfigReturns(infraConfig)

		hostConfig := &configfakes.FakeHostConfig{}
		infraConfig.GetCoordinatorReturns(hostConfig)
		hostConfig.GetIpReturns("192.12")

		err := testStageVldtr.ValidateConfig()

		testutils.AssertNotEmpty(t, err, "")
		testutils.Assert(t, "coordinator ip not valid IPv4", err.Error(), "")
	})

	t.Run("Validate config fails when deployment type is mirrored and standby not provided", func(t *testing.T) {

		testStageVldtr := validate.NewPrereqStageValidator()
		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}
		infraConfig := &configfakes.FakeInfraConfig{}
		testConfig.GetInfraConfigReturns(infraConfig)

		dbConfig := &configfakes.FakeDatabaseConfig{}
		testConfig.GetDatabaseConfigReturns(dbConfig)
		dbConfig.GetDeploymentTypeReturns(enums.DeploymentTypeMirrored)

		hostConfig := &configfakes.FakeHostConfig{}
		infraConfig.GetCoordinatorReturns(hostConfig)
		hostConfig.GetIpReturns("192.12.1.10")

		err := testStageVldtr.ValidateConfig()

		testutils.AssertNotEmpty(t, err, "")
		testutils.Assert(t, "standby config not found", err.Error(), "")
	})

	t.Run("Validate config fails when segmenthost config are not provided", func(t *testing.T) {

		testStageVldtr := validate.NewPrereqStageValidator()
		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}
		infraConfig := &configfakes.FakeInfraConfig{}
		testConfig.GetInfraConfigReturns(infraConfig)

		dbConfig := &configfakes.FakeDatabaseConfig{}
		testConfig.GetDatabaseConfigReturns(dbConfig)
		dbConfig.GetDeploymentTypeReturns(enums.DeploymentTypeMirrorless)

		hostConfig := &configfakes.FakeHostConfig{}
		infraConfig.GetCoordinatorReturns(hostConfig)
		hostConfig.GetIpReturns("192.12.1.11")

		err := testStageVldtr.ValidateConfig()

		testutils.AssertNotEmpty(t, err, "")
		testutils.Assert(t, "segmenthost config not found", err.Error(), "")
	})

	t.Run("Validate config fails when segmenthost ip list is not provided", func(t *testing.T) {

		testStageVldtr := validate.NewPrereqStageValidator()
		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}
		infraConfig := &configfakes.FakeInfraConfig{}
		testConfig.GetInfraConfigReturns(infraConfig)

		dbConfig := &configfakes.FakeDatabaseConfig{}
		testConfig.GetDatabaseConfigReturns(dbConfig)
		dbConfig.GetDeploymentTypeReturns(enums.DeploymentTypeMirrorless)

		hostConfig := &configfakes.FakeHostConfig{}
		infraConfig.GetCoordinatorReturns(hostConfig)
		hostConfig.GetIpReturns("192.12.1.13")

		err := testStageVldtr.ValidateConfig()

		testutils.AssertNotEmpty(t, err, "")
		testutils.Assert(t, "segmenthost config not found", err.Error(), "")
	})

	t.Run("Validate config fails when fail to generate segmentHost ip list", func(t *testing.T) {

		testStageVldtr := validate.NewPrereqStageValidator()
		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}
		infraConfig := &configfakes.FakeInfraConfig{}
		testConfig.GetInfraConfigReturns(infraConfig)

		dbConfig := &configfakes.FakeDatabaseConfig{}
		testConfig.GetDatabaseConfigReturns(dbConfig)
		dbConfig.GetDeploymentTypeReturns(enums.DeploymentTypeMirrorless)

		hostConfig := &configfakes.FakeHostConfig{}
		infraConfig.GetCoordinatorReturns(hostConfig)
		hostConfig.GetIpReturns("192.12.1.15")

		segHostConfig := &configfakes.FakeSegmentHostsConfig{}
		segHostConfig.GetNetworkReturns(&configfakes.FakeSegmentHostsNetworkConfig{})
		infraConfig.GetSegmentHostReturns(segHostConfig)

		testConfig.GenerateSegmentIPListReturns(fmt.Errorf("failed to generate segment ip list"))

		err := testStageVldtr.ValidateConfig()

		testutils.AssertNotEmpty(t, err, "")
		testutils.Assert(t, "failed to generate segment ip list", err.Error(), "")
	})

	t.Run("Validate config fails when artifact config is not present", func(t *testing.T) {

		testStageVldtr := validate.NewPrereqStageValidator()
		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}
		infraConfig := &configfakes.FakeInfraConfig{}
		testConfig.GetInfraConfigReturns(infraConfig)

		dbConfig := &configfakes.FakeDatabaseConfig{}
		testConfig.GetDatabaseConfigReturns(dbConfig)
		dbConfig.GetDeploymentTypeReturns(enums.DeploymentTypeMirrorless)

		hostConfig := &configfakes.FakeHostConfig{}
		infraConfig.GetCoordinatorReturns(hostConfig)
		hostConfig.GetIpReturns("192.12.1.1")

		segHostConfig := &configfakes.FakeSegmentHostsConfig{}
		segHostNetwork := &configfakes.FakeSegmentHostsNetworkConfig{}
		segHostConfig.GetNetworkReturns(segHostNetwork)
		infraConfig.GetSegmentHostReturns(segHostConfig)

		segHostNetwork.GetIpListReturns([]string{"192.12.1.2"})
		testConfig.GenerateSegmentIPListReturns(nil)

		err := testStageVldtr.ValidateConfig()

		testutils.AssertNotEmpty(t, err, "")
		testutils.Assert(t, "artifact config not found", err.Error(), "")
	})

	t.Run("Validate config succeeds for valid configurations", func(t *testing.T) {

		testStageVldtr := validate.NewPrereqStageValidator()
		testConfig := &configfakes.FakeConfig{}

		validate.Dependency.GetConfig = func() (config.Config, error) {
			return testConfig, nil
		}
		infraConfig := &configfakes.FakeInfraConfig{}
		testConfig.GetInfraConfigReturns(infraConfig)

		dbConfig := &configfakes.FakeDatabaseConfig{}
		testConfig.GetDatabaseConfigReturns(dbConfig)
		dbConfig.GetDeploymentTypeReturns(enums.DeploymentTypeMirrorless)

		hostConfig := &configfakes.FakeHostConfig{}
		infraConfig.GetCoordinatorReturns(hostConfig)
		hostConfig.GetIpReturns("192.12.1.1")

		segHostConfig := &configfakes.FakeSegmentHostsConfig{}
		segHostNetwork := &configfakes.FakeSegmentHostsNetworkConfig{}
		segHostConfig.GetNetworkReturns(segHostNetwork)
		infraConfig.GetSegmentHostReturns(segHostConfig)

		segHostNetwork.GetIpListReturns([]string{"192.12.1.2"})
		testConfig.GenerateSegmentIPListReturns(nil)

		testConfig.GetArtifactConfigReturns(&configfakes.FakeArtifactConfig{})

		err := testStageVldtr.ValidateConfig()

		testutils.Assert(t, nil, err, "")
	})
}
