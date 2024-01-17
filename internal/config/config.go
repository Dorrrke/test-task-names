package config

import (
	"flag"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	DbAddr  string `env:"DATA_BASE_ADDRES"`
	AppAddr string `env:"SERVER_ADDRES"`
	Env     string
}

func MustLoad() *Config {
	var cfg Config

	flag.StringVar(&cfg.AppAddr, "a", "", "addres and port to rin server")
	flag.StringVar(&cfg.DbAddr, "d", "", "data base addres")
	flag.StringVar(&cfg.Env, "env", "", "application env")
	flag.Parse()

	if cfg.DbAddr == "" {
		if err := env.Parse(&cfg.DbAddr); err != nil {
			panic("config value DbAdde not set, err: " + err.Error())
		}
	}
	if cfg.AppAddr == "" {
		if err := env.Parse(&cfg.AppAddr); err != nil {
			panic("config value AppAddr not set, err: " + err.Error())
		}
	}
	if cfg.Env == "" {
		cfg.Env = "local"
	}
	return &cfg
}
