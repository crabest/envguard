package cmd

import (
	"os"

	"github.com/crabest/envguard/internal/envmanager"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use <environment>",
	Short: "Use a specific environment",
	Long: `Use a specific environment by copying the specified environment
file from .envguard/ to the root .env file and tracking it as active.

This command automatically saves any changes to the current .env file
back to the currently active environment before switching.

Examples:
  envguard use production
  envguard use staging
  envguard use development`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envName := args[0]

		manager, err := envmanager.NewManager()
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		// Auto-sync current .env changes before switching
		manager.SyncActiveEnvironment()

		if err := manager.UseEnvironment(envName); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
