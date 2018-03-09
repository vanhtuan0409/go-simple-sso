package config

import "github.com/caarlos0/env"

type Config struct {
	SSS_URL    string `env:"SSO_URL"`
	SERVER_URL string `env:"SERVER_URL"`
	HTTP_PORT  int    `env:"HTTP_PORT"`
}

func Parse() *Config {
	cfg := new(Config)
	env.Parse(cfg)
	return cfg
}
