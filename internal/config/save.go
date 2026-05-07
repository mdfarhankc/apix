package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Save(config *Config) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}
