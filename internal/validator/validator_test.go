package validator

import (
	"testing"

	"github.com/crabest/envguard/internal/parser"
)

func TestValidateEnvFiles(t *testing.T) {
	envVars := parser.EnvVars{
		"DATABASE_URL": "test",
		"API_KEY":      "test",
		"PORT":         "3000",
		"EXTRA_VAR":    "extra",
	}

	exampleVars := parser.EnvVars{
		"DATABASE_URL": "example",
		"API_KEY":      "example",
		"PORT":         "8080",
		"MISSING_VAR":  "missing",
	}

	result := ValidateEnvFiles(envVars, exampleVars)

	expectedCommon := []string{"API_KEY", "DATABASE_URL", "PORT"}
	if len(result.CommonVars) != len(expectedCommon) {
		t.Errorf("Expected %d common variables, got %d", len(expectedCommon), len(result.CommonVars))
	}

	expectedMissing := []string{"MISSING_VAR"}
	if len(result.MissingVars) != len(expectedMissing) {
		t.Errorf("Expected %d missing variables, got %d", len(expectedMissing), len(result.MissingVars))
	}

	expectedExtra := []string{"EXTRA_VAR"}
	if len(result.ExtraVars) != len(expectedExtra) {
		t.Errorf("Expected %d extra variables, got %d", len(expectedExtra), len(result.ExtraVars))
	}

	if result.MissingVars[0] != "MISSING_VAR" {
		t.Errorf("Expected missing variable MISSING_VAR, got %s", result.MissingVars[0])
	}

	if result.ExtraVars[0] != "EXTRA_VAR" {
		t.Errorf("Expected extra variable EXTRA_VAR, got %s", result.ExtraVars[0])
	}
}

func TestValidateEnvFilesPerfectMatch(t *testing.T) {
	envVars := parser.EnvVars{
		"DATABASE_URL": "test",
		"API_KEY":      "test",
		"PORT":         "3000",
	}

	exampleVars := parser.EnvVars{
		"DATABASE_URL": "example",
		"API_KEY":      "example",
		"PORT":         "8080",
	}

	result := ValidateEnvFiles(envVars, exampleVars)

	if len(result.MissingVars) != 0 {
		t.Errorf("Expected no missing variables, got %d", len(result.MissingVars))
	}

	if len(result.ExtraVars) != 0 {
		t.Errorf("Expected no extra variables, got %d", len(result.ExtraVars))
	}

	if len(result.CommonVars) != 3 {
		t.Errorf("Expected 3 common variables, got %d", len(result.CommonVars))
	}
}

func TestValidateEnvFilesEmptyEnv(t *testing.T) {
	envVars := parser.EnvVars{}

	exampleVars := parser.EnvVars{
		"DATABASE_URL": "example",
		"API_KEY":      "example",
	}

	result := ValidateEnvFiles(envVars, exampleVars)

	if len(result.CommonVars) != 0 {
		t.Errorf("Expected no common variables, got %d", len(result.CommonVars))
	}

	if len(result.ExtraVars) != 0 {
		t.Errorf("Expected no extra variables, got %d", len(result.ExtraVars))
	}

	if len(result.MissingVars) != 2 {
		t.Errorf("Expected 2 missing variables, got %d", len(result.MissingVars))
	}
}

func TestValidateEnvFilesEmptyExample(t *testing.T) {
	envVars := parser.EnvVars{
		"DATABASE_URL": "test",
		"API_KEY":      "test",
	}

	exampleVars := parser.EnvVars{}

	result := ValidateEnvFiles(envVars, exampleVars)

	if len(result.CommonVars) != 0 {
		t.Errorf("Expected no common variables, got %d", len(result.CommonVars))
	}

	if len(result.MissingVars) != 0 {
		t.Errorf("Expected no missing variables, got %d", len(result.MissingVars))
	}

	if len(result.ExtraVars) != 2 {
		t.Errorf("Expected 2 extra variables, got %d", len(result.ExtraVars))
	}
}
