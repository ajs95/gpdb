package validate

import (
	"fmt"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type Stage string

const (
	StagePreReq    = Stage("pre-req")
	StageInfra     = Stage("infra")
	StageOsConfig  = Stage("os-config")
	StageGreenplum = Stage("greenplum")
)

//counterfeiter:generate . Validator
type Validator interface {
	Validate() error
}

type validateService struct{}

//counterfeiter:generate . Service
type Service interface {
	ValidateStage(Stage) error
}

func NewService() Service {
	return &validateService{}
}

func (*validateService) ValidateStage(stage Stage) error {

	stageValidator, err := Dependency.GetStageValidator(stage)
	if err != nil {
		return err
	}

	if err := stageValidator.ValidateConfig(); err != nil {
		return err
	}

	var errs []error
	for _, validator := range stageValidator.GetValidators() {
		if err := validator.Validate(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed for stage : %v, with errors: %v", stage, errs)
	}

	return nil
}

func getStageValidator(stage Stage) (StageValidator, error) {
	var stageValidator StageValidator
	switch stage {
	case StagePreReq:
		stageValidator = NewPrereqStageValidator()
	default:
		return nil, fmt.Errorf("invalid validation stage : %v", string(stage))
	}
	return stageValidator, nil
}

//counterfeiter:generate . StageValidator
type StageValidator interface {
	GetValidators() []Validator
	ValidateConfig() error
}
