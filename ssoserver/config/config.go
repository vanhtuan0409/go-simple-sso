package config

import "github.com/caarlos0/env"

type Config struct {
	DefaultCallback string `env:"DEFAULT_CALLBACK"`
	HTTP_PORT       int    `env:"HTTP_PORT"`
}

func Parse() *Config {
	cfg := new(Config)
	env.Parse(cfg)
	return cfg
}
