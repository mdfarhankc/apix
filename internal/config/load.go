package config

import (
	"encoding/json"
	"os"
)

func Load() (*Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				Environments: map[string]Environment{},
			}, nil
		}
		return nil, err
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	if config.Environments == nil {
		config.Environments = map[string]Environment{}
	}

	return &config, nil
}
