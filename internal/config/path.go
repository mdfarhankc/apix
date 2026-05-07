package config

import (
	"os"
	"path/filepath"
)

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".apix/config.json"), nil
}
