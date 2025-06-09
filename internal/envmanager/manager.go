package envmanager

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

const (
	EnvGuardDir = ".envguard"
	ActiveFile  = ".active"
)

type Manager struct {
	workingDir string
	envDir     string
}

func NewManager() (*Manager, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	envDir := filepath.Join(wd, EnvGuardDir)

	return &Manager{
		workingDir: wd,
		envDir:     envDir,
	}, nil
}

func (m *Manager) EnsureEnvGuardDir() error {
	if _, err := os.Stat(m.envDir); os.IsNotExist(err) {
		if err := os.MkdirAll(m.envDir, 0755); err != nil {
			return fmt.Errorf("failed to create %s directory: %w", EnvGuardDir, err)
		}
		color.Green("✅ Created %s directory", EnvGuardDir)
	}
	return nil
}

func (m *Manager) GetEnvPath(envName string) string {
	return filepath.Join(m.envDir, envName+".env")
}

func (m *Manager) GetRootEnvPath() string {
	return filepath.Join(m.workingDir, ".env")
}

func (m *Manager) GetActivePath() string {
	return filepath.Join(m.envDir, ActiveFile)
}

func (m *Manager) ListEnvironments() ([]string, error) {
	if err := m.EnsureEnvGuardDir(); err != nil {
		return nil, err
	}

	files, err := os.ReadDir(m.envDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s directory: %w", EnvGuardDir, err)
	}

	var envs []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".env") {
			envName := strings.TrimSuffix(file.Name(), ".env")
			envs = append(envs, envName)
		}
	}

	return envs, nil
}

func (m *Manager) EnvironmentExists(envName string) bool {
	envPath := m.GetEnvPath(envName)
	_, err := os.Stat(envPath)
	return err == nil
}

func (m *Manager) switchEnvironment(envName string) error {
	if err := m.EnsureEnvGuardDir(); err != nil {
		return err
	}

	envPath := m.GetEnvPath(envName)
	if !m.EnvironmentExists(envName) {
		return fmt.Errorf("environment '%s' does not exist in %s", envName, EnvGuardDir)
	}

	rootEnvPath := m.GetRootEnvPath()

	if err := m.copyFile(envPath, rootEnvPath); err != nil {
		return fmt.Errorf("failed to switch to environment '%s': %w", envName, err)
	}

	color.Blue("📁 Active .env file updated from %s/%s.env", EnvGuardDir, envName)

	return nil
}

func (m *Manager) CreateEnvironment(envName string, fromCurrent bool) error {
	if err := m.EnsureEnvGuardDir(); err != nil {
		return err
	}

	envPath := m.GetEnvPath(envName)

	if m.EnvironmentExists(envName) {
		return fmt.Errorf("environment '%s' already exists", envName)
	}

	var sourceFile string
	if fromCurrent {
		rootEnvPath := m.GetRootEnvPath()
		if _, err := os.Stat(rootEnvPath); os.IsNotExist(err) {
			color.Yellow("⚠️  No .env file found in root directory, creating empty environment")
			fromCurrent = false
		} else {
			sourceFile = rootEnvPath
		}
	}

	if fromCurrent && sourceFile != "" {
		if err := m.copyFile(sourceFile, envPath); err != nil {
			return fmt.Errorf("failed to copy current .env to '%s': %w", envName, err)
		}
		color.Green("✅ Created environment '%s' based on current .env", color.CyanString(envName))
	} else {
		file, err := os.Create(envPath)
		if err != nil {
			return fmt.Errorf("failed to create environment '%s': %w", envName, err)
		}
		file.Close()
		color.Green("✅ Created empty environment: %s", color.CyanString(envName))
	}

	color.Blue("📁 Environment file: %s/%s.env", EnvGuardDir, envName)
	return nil
}

func (m *Manager) DeleteEnvironment(envName string, confirm bool) error {
	if err := m.EnsureEnvGuardDir(); err != nil {
		return err
	}

	if !m.EnvironmentExists(envName) {
		return fmt.Errorf("environment '%s' does not exist", envName)
	}

	if confirm {
		color.Yellow("⚠️  Are you sure you want to delete environment '%s'? (y/N): ", envName)
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read user input: %w", err)
		}

		response = strings.ToLower(strings.TrimSpace(response))
		if response != "y" && response != "yes" {
			color.Blue("ℹ️  Deletion cancelled")
			return nil
		}
	}

	envPath := m.GetEnvPath(envName)
	if err := os.Remove(envPath); err != nil {
		return fmt.Errorf("failed to delete environment '%s': %w", envName, err)
	}

	color.Green("✅ Successfully deleted environment: %s", color.CyanString(envName))
	return nil
}

func (m *Manager) copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}

func (m *Manager) PromptForCurrentEnv() (bool, error) {
	rootEnvPath := m.GetRootEnvPath()
	if _, err := os.Stat(rootEnvPath); os.IsNotExist(err) {
		return false, nil
	}

	color.Yellow("🤔 Do you want to base the new environment on the current .env file? (Y/n): ")
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read user input: %w", err)
	}

	response = strings.ToLower(strings.TrimSpace(response))
	return response == "" || response == "y" || response == "yes", nil
}

func (m *Manager) SetActiveEnvironment(envName string) error {
	if err := m.EnsureEnvGuardDir(); err != nil {
		return err
	}

	activePath := m.GetActivePath()
	file, err := os.Create(activePath)
	if err != nil {
		return fmt.Errorf("failed to create active file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(envName)
	if err != nil {
		return fmt.Errorf("failed to write active environment: %w", err)
	}

	return nil
}

func (m *Manager) GetActiveEnvironment() (string, error) {
	if err := m.EnsureEnvGuardDir(); err != nil {
		return "", err
	}

	activePath := m.GetActivePath()
	if _, err := os.Stat(activePath); os.IsNotExist(err) {
		return "", fmt.Errorf("no active environment set")
	}

	content, err := os.ReadFile(activePath)
	if err != nil {
		return "", fmt.Errorf("failed to read active environment: %w", err)
	}

	envName := strings.TrimSpace(string(content))
	if envName == "" {
		return "", fmt.Errorf("active environment file is empty")
	}

	return envName, nil
}

func (m *Manager) UseEnvironment(envName string) error {
	if err := m.switchEnvironment(envName); err != nil {
		return err
	}

	if err := m.SetActiveEnvironment(envName); err != nil {
		return fmt.Errorf("failed to set active environment: %w", err)
	}

	color.Green("✅ Using environment: %s", color.CyanString(envName))
	return nil
}

// SyncActiveEnvironment automatically syncs changes from .env back to the active environment
func (m *Manager) SyncActiveEnvironment() error {
	activeEnv, err := m.GetActiveEnvironment()
	if err != nil {
		// No active environment set, nothing to sync
		return nil
	}

	rootEnvPath := m.GetRootEnvPath()
	envPath := m.GetEnvPath(activeEnv)

	// Check if .env file exists
	if _, err := os.Stat(rootEnvPath); os.IsNotExist(err) {
		// No .env file, nothing to sync
		return nil
	}

	// Check if stored environment file exists
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		// Environment file doesn't exist anymore, can't sync
		return nil
	}

	// Compare file contents to see if sync is needed
	if !m.filesAreDifferent(rootEnvPath, envPath) {
		// Files are the same, no sync needed
		return nil
	}

	// Copy .env changes back to the environment file
	if err := m.copyFile(rootEnvPath, envPath); err != nil {
		return fmt.Errorf("failed to sync .env changes to environment '%s': %w", activeEnv, err)
	}

	color.Blue("🔄 Synced .env changes to %s/%s.env", EnvGuardDir, activeEnv)
	return nil
}

// filesAreDifferent compares two files to see if they have different content
func (m *Manager) filesAreDifferent(file1, file2 string) bool {
	hash1, err1 := m.getFileHash(file1)
	hash2, err2 := m.getFileHash(file2)

	// If we can't read either file, assume they're different
	if err1 != nil || err2 != nil {
		return true
	}

	return hash1 != hash2
}

// getFileHash returns MD5 hash of file content
func (m *Manager) getFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
