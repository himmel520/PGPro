package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/himmel520/pgPro/internal/app"
	"github.com/himmel520/pgPro/internal/config"
	httphandler "github.com/himmel520/pgPro/internal/handler/http"
	"github.com/himmel520/pgPro/internal/repository/postgres"
	"github.com/himmel520/pgPro/internal/service"
	"github.com/sirupsen/logrus"
)

var (
	configPath = "./configs/base.yml"
)

func main() {
	// Load configuration
	cfg, err := config.New(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	// Initialize PostgreSQL repository
	repo, err := postgres.New(cfg)
	if err != nil {
		logrus.Fatalf("unable to connect to pool: %v", err)
	}
	defer repo.DB.Close()

	service := service.New(repo)
	handler := httphandler.New(service)

	// Start HTTP server
	srv := app.New(cfg.Server, handler.InitRoutes())
	go func() {
		logrus.Infof("the server is starting on :%v", cfg.Server.Port)

		if err := srv.Run(); err != nil {
			logrus.Errorf("error occured while running http server: %s", err.Error())
		}
	}()
	
	// Wait for termination signal
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	logrus.Info("the server is shutting down")

	// Shutdown server gracefully
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
