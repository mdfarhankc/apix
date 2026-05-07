package config

import (
	"fmt"
	"strings"
)

func ResolveURL(path string) (string, error) {
	if !strings.HasPrefix(path, "/") {
		return path, nil
	}

	cfg, err := Load()
	if err != nil {
		return "", err
	}

	if cfg.CurrentEnv == "" {
		return "", fmt.Errorf("no environment selected")
	}

	env, exists := cfg.Environments[cfg.CurrentEnv]
	if !exists {
		return "", fmt.Errorf("current environment not found")
	}

	return env.BaseURL + path, nil
}
