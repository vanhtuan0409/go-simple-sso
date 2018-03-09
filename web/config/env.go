package config

import "github.com/vanhtuan0409/go-simple-sso/web/service"

type AppEnv struct {
	Config             *Config
	TokenVerifyService service.TokenVerifyService
}
