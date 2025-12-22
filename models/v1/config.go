package v1

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Auth struct {
	GithubApi string `toml:"github_api"`
}

type Config struct {
	Version uint  `toml:"version"`
	Auth    *Auth `toml:"auth"`
}

func ReadFromConfigToml(filename string) (*Config, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := toml.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadEnvironmentVariable(config *Config) *Config {
	v, ok := os.LookupEnv("GITHUB_API")
	if !ok {
		return config
	}

	config.Auth.GithubApi = v

	return config
}
