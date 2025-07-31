package app

import (
	"context"
	"log/slog"
	"os"
	"syscall"
	"os/signal"

	"github.com/2pizzzza/cryptosync/internal/config"
	"github.com/2pizzzza/cryptosync/pkg/httpserver"
	"github.com/2pizzzza/cryptosync/pkg/logger"
	"github.com/2pizzzza/cryptosync/pkg/postgres"
)

func New(cfg *config.Config){
	ctx := context.Background()
	log := logger.New(cfg.Log.Level)

	application := httpserver.New(log, cfg)
	log.Info("Server is live")
	conn, err := postgres.New(ctx, cfg)
	_ = conn
	if err != nil {
		log.Error("Failed connect to database")
		application.Stop()
	}

		go application.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal:", sign.String()))

	application.Stop()

	log.Info("Server is dead")

}