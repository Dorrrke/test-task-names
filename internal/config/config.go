package config

import (
	"flag"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	DbAddr  DBAddrEnv
	AppAddr AppAddrEnv
	Env     string
}

type DBAddrEnv struct {
	DbAddr string `env:"DATA_BASE_ADDRES,required"`
}
type AppAddrEnv struct {
	AppAddr string `env:"SERVER_ADDRES,required"`
}

func MustLoad() *Config {
	var cfg Config

	flag.StringVar(&cfg.AppAddr.AppAddr, "a", "", "addres and port to rin server")
	flag.StringVar(&cfg.DbAddr.DbAddr, "d", "", "data base addres")
	flag.StringVar(&cfg.Env, "env", "", "application env")
	flag.Parse()

	if cfg.DbAddr.DbAddr == "" {
		if err := env.Parse(&cfg.DbAddr); err != nil {
			panic("config value DbAddr not set, err: " + err.Error())
		}
	}
	if cfg.AppAddr.AppAddr == "" {
		if err := env.Parse(&cfg.AppAddr); err != nil {
			panic("config value AppAddr not set, err: " + err.Error())
		}
	}
	if cfg.Env == "" {
		cfg.Env = "local"
	}
	return &cfg
}
