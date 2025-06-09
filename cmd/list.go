package cmd

import (
	"fmt"
	"os"

	"github.com/crabest/envguard/internal/envmanager"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available environments",
	Long: `List all environment files stored in the .envguard/ directory.
Shows the environment names that can be used with use, delete, etc.

Example:
  envguard list`,
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := envmanager.NewManager()
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		// Auto-sync .env changes to active environment
		manager.SyncActiveEnvironment()

		environments, err := manager.ListEnvironments()
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		color.Cyan("🌍 Available Environments:")
		color.Cyan("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

		if len(environments) == 0 {
			color.Yellow("📭 No environments found in .envguard/ directory")
			color.Blue("💡 Create your first environment with: envguard create -e <n>")
			return
		}

		for i, env := range environments {
			status := fmt.Sprintf("%d.", i+1)
			fmt.Printf("   %s %s\n", color.BlueString(status), color.GreenString(env))
		}

		fmt.Printf("\n📊 Total: %s environments\n", color.CyanString(fmt.Sprintf("%d", len(environments))))
		color.Blue("💡 Use with: envguard use <environment>")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
