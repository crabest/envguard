package cmd

import (
	"fmt"
	"os"

	"github.com/crabest/envguard/internal/envmanager"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the currently active environment",
	Long: `Show the currently active environment by reading from .envguard/.active.
This shows which environment was last activated using 'envguard use <env>'.

Example:
  envguard status`,
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := envmanager.NewManager()
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		// Auto-sync .env changes to active environment
		manager.SyncActiveEnvironment()

		activeEnv, err := manager.GetActiveEnvironment()
		if err != nil {
			color.Yellow("⚠️  %v", err)
			color.Blue("💡 Use 'envguard use <environment>' to set an active environment")
			return
		}

		// Check if the active environment file still exists
		if !manager.EnvironmentExists(activeEnv) {
			color.Red("❌ Active environment '%s' no longer exists", activeEnv)
			color.Blue("💡 Available environments:")

			environments, listErr := manager.ListEnvironments()
			if listErr == nil {
				for i, env := range environments {
					fmt.Printf("   %d. %s\n", i+1, color.GreenString(env))
				}
			}
			color.Blue("💡 Use 'envguard use <environment>' to set a new active environment")
			return
		}

		color.Cyan("🌍 Environment Status:")
		color.Cyan("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

		fmt.Printf("📍 Active Environment: %s\n", color.GreenString(activeEnv))
		fmt.Printf("📁 Environment File: %s/%s.env\n",
			color.BlueString(".envguard"),
			color.BlueString(activeEnv))
		fmt.Printf("🎯 Active .env: %s\n", color.BlueString(".env"))

		color.Green("✅ Environment '%s' is currently active", activeEnv)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
