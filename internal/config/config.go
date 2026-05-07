package config

type Environment struct {
	BaseURL string `json:"base_url"`
}

type Config struct {
	CurrentEnv   string                 `json:"current_env"`
	Environments map[string]Environment `json:"environments"`
}
