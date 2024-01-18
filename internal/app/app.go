package app

import (
	"context"
	"net/http"

	"github.com/Dorrrke/test-task-names/internal/logger"
	"github.com/Dorrrke/test-task-names/internal/server"
	"go.uber.org/zap"
)

type App struct {
	httpServer *server.Server
}

func New(
	addr string,
	storage server.Storage,
	enrichService server.Enrichment,
) *App {
	serv := http.Server{
		Addr: addr,
	}
	server := server.New(&serv, storage, enrichService)
	server.RegisterServer()
	return &App{
		httpServer: server,
	}
}

func (a *App) Run() error {
	logger.Log.Info("Start http server", zap.String("addr", a.httpServer.HttpServer.Addr))
	return a.httpServer.HttpServer.ListenAndServe()
}

func (a *App) Stop() error {
	logger.Log.Info("Stop http server")
	return a.httpServer.HttpServer.Shutdown(context.Background())
}
