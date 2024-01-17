package logger

import (
	"go.uber.org/zap"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var Log *zap.Logger = zap.NewNop()

func SetupLogger(env string) error {
	cfg := zap.NewProductionConfig()
	switch env {
	case envLocal:
		lvl, err := zap.ParseAtomicLevel("debug")
		if err != nil {
			return err
		}
		cfg.Level = lvl
		cfg.Encoding = "console"

		zl, err := cfg.Build()
		if err != nil {
			return err
		}
		Log = zl
	case envDev:
		lvl, err := zap.ParseAtomicLevel("debug")
		if err != nil {
			return err
		}
		cfg.Level = lvl
		cfg.Encoding = "json"

		zl, err := cfg.Build()
		if err != nil {
			return err
		}

		Log = zl
	case envProd:
		lvl, err := zap.ParseAtomicLevel("info")
		if err != nil {
			return err
		}
		cfg.Level = lvl

		cfg.Encoding = "json"
		cfg.OutputPaths = []string{"stdout", "/tmp/logs"}
		cfg.ErrorOutputPaths = []string{"stderr"}

		zl, err := cfg.Build()
		if err != nil {
			return err
		}

		Log = zl
	}

	return nil
}
