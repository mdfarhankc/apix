package config

import (
	"fmt"
	"strings"
)

type Resolved struct {
	URL     string
	Headers map[string]string
}

func Resolve(path string) (*Resolved, error) {
	if !strings.HasPrefix(path, "/") {
		return &Resolved{URL: path, Headers: map[string]string{}}, nil
	}

	cfg, err := Load()
	if err != nil {
		return nil, err
	}

	if cfg.CurrentEnv == "" {
		return nil, fmt.Errorf("no environment selected")
	}

	env, exists := cfg.Environments[cfg.CurrentEnv]
	if !exists {
		return nil, fmt.Errorf("current environment not found")
	}

	return &Resolved{
		URL:     env.BaseURL + path,
		Headers: env.AuthHeaders(),
	}, nil
}
