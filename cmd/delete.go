package cmd

import (
	"os"

	"github.com/crabest/envguard/internal/envmanager"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an environment",
	Long: `Delete an environment file from the .envguard/ directory.
By default, this command will ask for confirmation before deleting.

Examples:
  envguard delete -e staging
  envguard delete --env old-config
  envguard delete -e test --no-confirm`,
	Run: func(cmd *cobra.Command, args []string) {
		envName, _ := cmd.Flags().GetString("env")
		if envName == "" {
			color.Red("Error: environment name is required")
			color.Yellow("Usage: envguard delete -e <environment>")
			os.Exit(1)
		}

		noConfirm, _ := cmd.Flags().GetBool("no-confirm")
		confirm := !noConfirm

		manager, err := envmanager.NewManager()
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		if err := manager.DeleteEnvironment(envName, confirm); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	deleteCmd.Flags().StringP("env", "e", "", "Environment name to delete (required)")
	deleteCmd.Flags().Bool("no-confirm", false, "Skip confirmation prompt")
	deleteCmd.MarkFlagRequired("env")
	rootCmd.AddCommand(deleteCmd)
}
