package parser

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVars map[string]string

func ParseEnvFile(filename string) (EnvVars, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}

	vars, err := godotenv.Read(filename)
	if err != nil {
		return nil, err
	}

	envVars := make(EnvVars)
	for key, value := range vars {
		envVars[key] = value
	}

	return envVars, nil
}

func GetVariableNames(envVars EnvVars) []string {
	var names []string
	for name := range envVars {
		names = append(names, name)
	}
	return names
}

func HasVariable(envVars EnvVars, name string) bool {
	_, exists := envVars[name]
	return exists
}
