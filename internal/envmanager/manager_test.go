package envmanager

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewManager(t *testing.T) {
	manager, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	if manager.workingDir == "" {
		t.Error("Working directory should not be empty")
	}

	if manager.envDir == "" {
		t.Error("Env directory should not be empty")
	}

	expectedEnvDir := filepath.Join(manager.workingDir, EnvGuardDir)
	if manager.envDir != expectedEnvDir {
		t.Errorf("Expected env dir %s, got %s", expectedEnvDir, manager.envDir)
	}
}

func TestEnvironmentExists(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "envguard-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tmpDir)

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	if manager.EnvironmentExists("nonexistent") {
		t.Error("Environment should not exist")
	}

	if err := manager.EnsureEnvGuardDir(); err != nil {
		t.Fatalf("Failed to create envguard dir: %v", err)
	}

	envPath := manager.GetEnvPath("test")
	file, err := os.Create(envPath)
	if err != nil {
		t.Fatalf("Failed to create test env file: %v", err)
	}
	file.Close()

	if !manager.EnvironmentExists("test") {
		t.Error("Environment should exist")
	}
}

func TestListEnvironments(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "envguard-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tmpDir)

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	envs, err := manager.ListEnvironments()
	if err != nil {
		t.Fatalf("Failed to list environments: %v", err)
	}

	if len(envs) != 0 {
		t.Errorf("Expected 0 environments, got %d", len(envs))
	}

	if err := manager.CreateEnvironment("test1", false); err != nil {
		t.Fatalf("Failed to create environment: %v", err)
	}

	if err := manager.CreateEnvironment("test2", false); err != nil {
		t.Fatalf("Failed to create environment: %v", err)
	}

	envs, err = manager.ListEnvironments()
	if err != nil {
		t.Fatalf("Failed to list environments: %v", err)
	}

	if len(envs) != 2 {
		t.Errorf("Expected 2 environments, got %d", len(envs))
	}

	expectedEnvs := map[string]bool{"test1": true, "test2": true}
	for _, env := range envs {
		if !expectedEnvs[env] {
			t.Errorf("Unexpected environment: %s", env)
		}
	}
}

func TestCreateEnvironment(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "envguard-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tmpDir)

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	if err := manager.CreateEnvironment("test", false); err != nil {
		t.Fatalf("Failed to create environment: %v", err)
	}

	if !manager.EnvironmentExists("test") {
		t.Error("Environment should exist after creation")
	}

	if err := manager.CreateEnvironment("test", false); err == nil {
		t.Error("Should not be able to create duplicate environment")
	}
}

func TestUseEnvironmentSwitching(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "envguard-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tmpDir)

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	if err := manager.CreateEnvironment("test", false); err != nil {
		t.Fatalf("Failed to create environment: %v", err)
	}

	envPath := manager.GetEnvPath("test")
	file, err := os.OpenFile(envPath, os.O_WRONLY, 0644)
	if err != nil {
		t.Fatalf("Failed to open env file: %v", err)
	}
	file.WriteString("TEST_VAR=test_value\n")
	file.Close()

	if err := manager.UseEnvironment("test"); err != nil {
		t.Fatalf("Failed to use environment: %v", err)
	}

	rootEnvPath := manager.GetRootEnvPath()
	content, err := os.ReadFile(rootEnvPath)
	if err != nil {
		t.Fatalf("Failed to read root env file: %v", err)
	}

	expectedContent := "TEST_VAR=test_value\n"
	if string(content) != expectedContent {
		t.Errorf("Expected content %q, got %q", expectedContent, string(content))
	}

	// Also check that active environment was set
	activeEnv, err := manager.GetActiveEnvironment()
	if err != nil {
		t.Fatalf("Failed to get active environment: %v", err)
	}

	if activeEnv != "test" {
		t.Errorf("Expected active environment 'test', got '%s'", activeEnv)
	}
}

func TestDeleteEnvironment(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "envguard-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tmpDir)

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	if err := manager.CreateEnvironment("test", false); err != nil {
		t.Fatalf("Failed to create environment: %v", err)
	}

	if !manager.EnvironmentExists("test") {
		t.Error("Environment should exist before deletion")
	}

	if err := manager.DeleteEnvironment("test", false); err != nil {
		t.Fatalf("Failed to delete environment: %v", err)
	}

	if manager.EnvironmentExists("test") {
		t.Error("Environment should not exist after deletion")
	}

	if err := manager.DeleteEnvironment("nonexistent", false); err == nil {
		t.Error("Should not be able to delete non-existent environment")
	}
}

func TestSetAndGetActiveEnvironment(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "envguard-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tmpDir)

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	// Test getting active environment when none is set
	_, err = manager.GetActiveEnvironment()
	if err == nil {
		t.Error("Should return error when no active environment is set")
	}

	// Test setting active environment
	if err := manager.SetActiveEnvironment("test"); err != nil {
		t.Fatalf("Failed to set active environment: %v", err)
	}

	// Test getting active environment
	activeEnv, err := manager.GetActiveEnvironment()
	if err != nil {
		t.Fatalf("Failed to get active environment: %v", err)
	}

	if activeEnv != "test" {
		t.Errorf("Expected active environment 'test', got '%s'", activeEnv)
	}

	// Test changing active environment
	if err := manager.SetActiveEnvironment("production"); err != nil {
		t.Fatalf("Failed to set active environment: %v", err)
	}

	activeEnv, err = manager.GetActiveEnvironment()
	if err != nil {
		t.Fatalf("Failed to get active environment: %v", err)
	}

	if activeEnv != "production" {
		t.Errorf("Expected active environment 'production', got '%s'", activeEnv)
	}
}

func TestUseEnvironment(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "envguard-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tmpDir)

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	// Create a test environment
	if err := manager.CreateEnvironment("test", false); err != nil {
		t.Fatalf("Failed to create environment: %v", err)
	}

	// Add content to the environment file
	envPath := manager.GetEnvPath("test")
	file, err := os.OpenFile(envPath, os.O_WRONLY, 0644)
	if err != nil {
		t.Fatalf("Failed to open env file: %v", err)
	}
	file.WriteString("TEST_VAR=test_value\n")
	file.Close()

	// Use the environment
	if err := manager.UseEnvironment("test"); err != nil {
		t.Fatalf("Failed to use environment: %v", err)
	}

	// Check that the root .env file was updated
	rootEnvPath := manager.GetRootEnvPath()
	content, err := os.ReadFile(rootEnvPath)
	if err != nil {
		t.Fatalf("Failed to read root env file: %v", err)
	}

	expectedContent := "TEST_VAR=test_value\n"
	if string(content) != expectedContent {
		t.Errorf("Expected content %q, got %q", expectedContent, string(content))
	}

	// Check that the active environment was set
	activeEnv, err := manager.GetActiveEnvironment()
	if err != nil {
		t.Fatalf("Failed to get active environment: %v", err)
	}

	if activeEnv != "test" {
		t.Errorf("Expected active environment 'test', got '%s'", activeEnv)
	}
}
