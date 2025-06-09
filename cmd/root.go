package cmd

import (
	"fmt"
	"os"

	"github.com/crabest/envguard/internal/envmanager"
	"github.com/crabest/envguard/internal/parser"
	"github.com/crabest/envguard/internal/validator"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	envFile     string
	exampleFile string
)

var rootCmd = &cobra.Command{
	Use:   "envguard",
	Short: "Validate .env files and manage multiple environments",
	Long: `EnvGuard is a CLI tool that validates your .env files against .env.example files
and manages multiple environment configurations using a hidden .envguard/ directory.

Features:
• Validate .env files against .env.example templates
• Manage multiple environments (use, create, list, delete)
• Colored output with detailed summaries
• Environment-specific configuration management

Examples:
  envguard                        # Validate current .env against .env.example
  envguard use production         # Use production environment
  envguard create -e staging      # Create new staging environment
  envguard list                   # List all available environments`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runValidation(); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&envFile, "env", "e", ".env", "Path to the .env file")
	rootCmd.Flags().StringVarP(&exampleFile, "example", "x", ".env.example", "Path to the .env.example file")
}

func Execute() error {
	return rootCmd.Execute()
}

func runValidation() error {
	// Auto-sync .env changes to active environment before validation
	manager, err := envmanager.NewManager()
	if err == nil {
		// Ignore sync errors for now, validation should still proceed
		manager.SyncActiveEnvironment()
	}

	color.Cyan("🔍 EnvGuard - Environment File Validator\n")

	envVars, err := parser.ParseEnvFile(envFile)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", envFile, err)
	}

	exampleVars, err := parser.ParseEnvFile(exampleFile)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", exampleFile, err)
	}

	result := validator.ValidateEnvFiles(envVars, exampleVars)

	validator.PrintResults(result, envFile, exampleFile)

	if len(result.MissingVars) > 0 {
		return fmt.Errorf("validation failed: %d missing variables", len(result.MissingVars))
	}

	return nil
}
