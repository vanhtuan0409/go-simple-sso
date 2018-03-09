package config

import "github.com/caarlos0/env"

type Config struct {
	SSO_URL    string `env:"SSO_URL"`
	SERVER_URL string `env:"SERVER_URL"`
	HTTP_PORT  int    `env:"HTTP_PORT" envDefault:"8080"`
	APP_TITLE  string `env:"APP_TITLE" envDefault:"Web Client"`
}

func Parse() *Config {
	cfg := new(Config)
	env.Parse(cfg)
	return cfg
}
