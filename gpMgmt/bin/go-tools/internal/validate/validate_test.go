package validate_test

import (
	"fmt"
	"testing"

	"github.com/greenplum-db/gpdb/gp/internal/validate"
	"github.com/greenplum-db/gpdb/gp/internal/validate/validatefakes"
	"github.com/greenplum-db/gpdb/gp/testutils"
	"github.com/greenplum-db/gpdb/gp/testutils/exectest"
)

func init() {
	exectest.RegisterMains()
}

func TestValidateStage(t *testing.T) {

	t.Run("ValidateStage fails for invalid stage", func(t *testing.T) {

		testValidateService := validate.NewService()

		err := testValidateService.ValidateStage(validate.Stage("some_stage"))

		testutils.Assert(t, "invalid validation stage : some_stage", err.Error(), "")
	})

	t.Run("ValidateStage succeeds for valid stage", func(t *testing.T) {

		validate.Dependency.GetStageValidator = func(s validate.Stage) (validate.StageValidator, error) {
			testStageValidator := &validatefakes.FakeStageValidator{}
			testStageValidator.GetValidatorsReturns([]validate.Validator{})
			testStageValidator.ValidateConfigReturns(fmt.Errorf("config validation failed for stage pre-req"))
			return testStageValidator, nil
		}

		validStages := []validate.Stage{
			validate.StagePreReq,
		}

		for _, stage := range validStages {
			testValidateService := validate.NewService()
			err := testValidateService.ValidateStage(stage)
			testutils.AssertNotEmpty(t, err, "")
			testutils.Assert(t, "config validation failed for stage pre-req", err.Error(), "")
		}
	})

	t.Run("ValidateStage succeeds for valid stage", func(t *testing.T) {

		validate.Dependency.GetStageValidator = func(s validate.Stage) (validate.StageValidator, error) {
			testStageValidator := &validatefakes.FakeStageValidator{}
			testStageValidator.GetValidatorsReturns([]validate.Validator{})
			testStageValidator.ValidateConfigReturns(nil)
			return testStageValidator, nil
		}

		validStages := []validate.Stage{
			validate.StagePreReq,
		}

		for _, stage := range validStages {
			testValidateService := validate.NewService()
			err := testValidateService.ValidateStage(stage)
			testutils.Assert(t, nil, err, "")
		}
	})
}
