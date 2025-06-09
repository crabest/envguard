package cmd

import (
	"os"

	"github.com/crabest/envguard/internal/envmanager"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new environment",
	Long: `Create a new environment file in the .envguard/ directory.
You can create an empty environment or base it on the current .env file.

Examples:
  envguard create -e staging
  envguard create -e production --from-current
  envguard create --env development`,
	Run: func(cmd *cobra.Command, args []string) {
		envName, _ := cmd.Flags().GetString("env")
		if envName == "" {
			color.Red("Error: environment name is required")
			color.Yellow("Usage: envguard create -e <environment>")
			os.Exit(1)
		}

		fromCurrent, _ := cmd.Flags().GetBool("from-current")

		manager, err := envmanager.NewManager()
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		// Auto-sync .env changes to active environment before creating new one
		manager.SyncActiveEnvironment()

		if !fromCurrent {
			shouldBase, err := manager.PromptForCurrentEnv()
			if err != nil {
				color.Red("Error: %v", err)
				os.Exit(1)
			}
			fromCurrent = shouldBase
		}

		if err := manager.CreateEnvironment(envName, fromCurrent); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.Flags().StringP("env", "e", "", "Environment name to create (required)")
	createCmd.Flags().Bool("from-current", false, "Base new environment on current .env file")
	createCmd.MarkFlagRequired("env")
	rootCmd.AddCommand(createCmd)
}
