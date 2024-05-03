package validate

import "fmt"

type artifactValidator struct{}

func NewArtifactValidator() Validator {
	return &artifactValidator{}
}

func (artifactValidator) Validate() error {
	conf, err := Dependency.GetConfig()
	if err != nil {
		return err
	}

	fsUtil := Dependency.NewFileSystem()

	if fileFound, err := fsUtil.IsFilePresent(conf.GetArtifactConfig().GetGreenplum()); err != nil {
		return err
	} else if !fileFound {
		return fmt.Errorf("artifact not found : %s", conf.GetArtifactConfig().GetGreenplum())
	}

	for _, artifact := range conf.GetArtifactConfig().GetDependencyList() {
		if fileFound, err := fsUtil.IsFilePresent(artifact); err != nil {
			return err
		} else if !fileFound {
			return fmt.Errorf("artifact not found : %s", artifact)
		}
	}

	return nil
}
