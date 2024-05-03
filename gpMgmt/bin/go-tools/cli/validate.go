package cli

import (
	"os"

	"github.com/greenplum-db/gpdb/gp/internal/config"
	"github.com/greenplum-db/gpdb/gp/internal/validate"
	"github.com/spf13/cobra"
)

var appConfig config.Config
var initStage string

// validateCmd
func validateCmd() *cobra.Command {
	validateCmd := &cobra.Command{
		Use:     "validate",
		Short:   "Validate cluster, segments",
		PreRunE: InitCommand,
		RunE:    RunValidateCmd,
	}

	validateCmd.Flags().StringVarP(&initStage, "stage", "s", "", "Stage (required)")
	validateCmd.MarkFlagRequired("stage")
	return validateCmd
}

func InitCommand(cmd *cobra.Command, args []string) error {
	appConfig = config.New()
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return appConfig.Load(home)
}

func RunValidateCmd(cmd *cobra.Command, args []string) error {
	validateSvc := validate.NewService()
	return validateSvc.ValidateStage(validate.Stage(initStage))
}
