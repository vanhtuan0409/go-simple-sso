package config

import "github.com/vanhtuan0409/go-simple-sso/ssoserver/datastore"

type AppEnv struct {
	Config    *Config
	Datastore datastore.Datastore
}
