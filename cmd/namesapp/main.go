package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dorrrke/test-task-names/internal/app"
	"github.com/Dorrrke/test-task-names/internal/config"
	"github.com/Dorrrke/test-task-names/internal/logger"
	"github.com/Dorrrke/test-task-names/internal/services"
	"github.com/Dorrrke/test-task-names/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load("values.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

		<-c
		cancel()
	}()
	//TODO: Инициализация конфига
	cfg := config.MustLoad()

	//TODO: Инициализация логгера
	if err := logger.SetupLogger(cfg.Env); err != nil {
		panic("Setup logger error: " + err.Error())
	}

	pool, err := pgxpool.New(ctx, cfg.DbAddr.DbAddr)
	if err != nil {
		logger.Log.Error("database connect error", zap.Error(err))
		panic(err)
	}
	defer pool.Close()

	storage := storage.New(pool)

	enrichService := services.New()

	logger.Log.Info("starting application")
	logger.Log.Debug("application config",
		zap.String("server addr", cfg.AppAddr.AppAddr),
		zap.String("db addr", cfg.DbAddr.DbAddr),
		zap.String("env", cfg.Env))

	//TODO: Инициализация приложения
	application := app.New(cfg.AppAddr.AppAddr, storage, enrichService)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return application.Run()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return application.Stop()
	})

	if err := g.Wait(); err != nil {
		logger.Log.Error("exit reason:", zap.Error(err))
	}

	//TODO: инит сервиса и харнилища

}
