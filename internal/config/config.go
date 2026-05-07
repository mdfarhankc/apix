package config

type Environment struct {
	BaseURL string `json:"base_url"`

	BearerToken string `json:"bearer_token,omitempty"`
}

func (e Environment) AuthHeaders() map[string]string {
	h := map[string]string{}
	if e.BearerToken != "" {
		h["Authorization"] = "Bearer " + e.BearerToken
	}
	return h
}

type Config struct {
	CurrentEnv   string                 `json:"current_env"`
	Environments map[string]Environment `json:"environments"`
}
