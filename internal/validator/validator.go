package validator

import (
	"fmt"
	"sort"

	"github.com/crabest/envguard/internal/parser"

	"github.com/fatih/color"
)

type ValidationResult struct {
	MissingVars []string
	ExtraVars   []string
	CommonVars  []string
}

func ValidateEnvFiles(envVars, exampleVars parser.EnvVars) ValidationResult {
	result := ValidationResult{
		MissingVars: []string{},
		ExtraVars:   []string{},
		CommonVars:  []string{},
	}

	exampleNames := parser.GetVariableNames(exampleVars)
	envNames := parser.GetVariableNames(envVars)

	for _, name := range exampleNames {
		if parser.HasVariable(envVars, name) {
			result.CommonVars = append(result.CommonVars, name)
		} else {
			result.MissingVars = append(result.MissingVars, name)
		}
	}

	for _, name := range envNames {
		if !parser.HasVariable(exampleVars, name) {
			result.ExtraVars = append(result.ExtraVars, name)
		}
	}

	sort.Strings(result.MissingVars)
	sort.Strings(result.ExtraVars)
	sort.Strings(result.CommonVars)

	return result
}

func PrintResults(result ValidationResult, envFile, exampleFile string) {
	color.Cyan("\n📊 Validation Results:")
	color.Cyan("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	fmt.Printf("📁 Comparing: %s ↔ %s\n\n", color.BlueString(envFile), color.BlueString(exampleFile))

	if len(result.CommonVars) > 0 {
		color.Green("✅ Variables found in both files (%d):", len(result.CommonVars))
		for _, name := range result.CommonVars {
			fmt.Printf("   ✓ %s\n", color.GreenString(name))
		}
		fmt.Println()
	}

	if len(result.MissingVars) > 0 {
		color.Yellow("⚠️  Missing variables in %s (%d):", envFile, len(result.MissingVars))
		for _, name := range result.MissingVars {
			fmt.Printf("   • %s\n", color.YellowString(name))
		}
		fmt.Println()
	}

	if len(result.ExtraVars) > 0 {
		color.Red("❌ Extra variables in %s not found in %s (%d):", envFile, exampleFile, len(result.ExtraVars))
		for _, name := range result.ExtraVars {
			fmt.Printf("   • %s\n", color.RedString(name))
		}
		fmt.Println()
	}

	PrintSummary(result)
}

func PrintSummary(result ValidationResult) {
	color.Cyan("📈 Summary:")
	color.Cyan("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	okCount := len(result.CommonVars)
	missingCount := len(result.MissingVars)
	extraCount := len(result.ExtraVars)

	var status string
	if missingCount == 0 && extraCount == 0 {
		status = color.GreenString("🎉 Perfect! All environment variables are properly configured.")
	} else if missingCount > 0 && extraCount == 0 {
		status = color.YellowString("⚠️  Some variables are missing from your .env file.")
	} else if missingCount == 0 && extraCount > 0 {
		status = color.BlueString("ℹ️  You have extra variables in your .env file.")
	} else {
		status = color.RedString("❌ Your .env file has missing and extra variables.")
	}

	fmt.Printf("%s\n\n", status)

	fmt.Printf("📊 %s %d variables OK",
		color.GreenString("✅"), okCount)

	if missingCount > 0 {
		fmt.Printf(" • %s %d missing",
			color.YellowString("⚠️"), missingCount)
	}

	if extraCount > 0 {
		fmt.Printf(" • %s %d unused",
			color.RedString("❌"), extraCount)
	}

	fmt.Println()
}
