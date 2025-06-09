package parser

import (
	"os"
	"testing"
)

func TestParseEnvFile(t *testing.T) {
	testContent := `# Test env file
DATABASE_URL=postgresql://localhost:5432/test
API_KEY=test-key
PORT=3000
DEBUG=true
`

	tmpFile, err := os.CreateTemp("", "test.env")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(testContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	envVars, err := ParseEnvFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to parse env file: %v", err)
	}

	expected := map[string]string{
		"DATABASE_URL": "postgresql://localhost:5432/test",
		"API_KEY":      "test-key",
		"PORT":         "3000",
		"DEBUG":        "true",
	}

	if len(envVars) != len(expected) {
		t.Errorf("Expected %d variables, got %d", len(expected), len(envVars))
	}

	for key, value := range expected {
		if envVars[key] != value {
			t.Errorf("Expected %s=%s, got %s=%s", key, value, key, envVars[key])
		}
	}
}

func TestParseEnvFileNotFound(t *testing.T) {
	_, err := ParseEnvFile("nonexistent.env")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestGetVariableNames(t *testing.T) {
	envVars := EnvVars{
		"DATABASE_URL": "test",
		"API_KEY":      "test",
		"PORT":         "3000",
	}

	names := GetVariableNames(envVars)
	if len(names) != 3 {
		t.Errorf("Expected 3 variable names, got %d", len(names))
	}

	expectedNames := map[string]bool{
		"DATABASE_URL": true,
		"API_KEY":      true,
		"PORT":         true,
	}

	for _, name := range names {
		if !expectedNames[name] {
			t.Errorf("Unexpected variable name: %s", name)
		}
	}
}

func TestHasVariable(t *testing.T) {
	envVars := EnvVars{
		"DATABASE_URL": "test",
		"API_KEY":      "test",
	}

	if !HasVariable(envVars, "DATABASE_URL") {
		t.Error("Expected DATABASE_URL to exist")
	}

	if !HasVariable(envVars, "API_KEY") {
		t.Error("Expected API_KEY to exist")
	}

	if HasVariable(envVars, "NONEXISTENT") {
		t.Error("Expected NONEXISTENT to not exist")
	}
}
