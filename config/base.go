package config

type BaseConfig struct {
	CurrentEnv   string                 `yaml:"current_env"`
	Identity     Identity               `yaml:"identity"`
	Environments map[string]Environment `yaml:"environments"`
}

type Identity struct {
	Firebase FirebaseIdentity `yaml:"firebase"`
}

type FirebaseIdentity struct {
	URL string `yaml:"url"`
}

type Environment struct {
	BaseURL string `yaml:"base_url"`
}
