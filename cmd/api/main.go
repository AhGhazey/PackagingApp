package main

import (
	"context"
	"github/ahmedghazey/packaging/internal/configuration"
	"github/ahmedghazey/packaging/internal/http/handler"
	"github/ahmedghazey/packaging/internal/server"
	"github/ahmedghazey/packaging/internal/service"
	"github/ahmedghazey/packaging/internal/storage/inmemory"
	"github/ahmedghazey/packaging/pkg/logging"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := configuration.GetConfiguration()
	if err != nil {
		log.Fatal("unable to read configuration", err)
	}
	err = logging.InitLogger(config.LogLevel, config.ServiceName, config.Environment)
	if err != nil {
		log.Fatal("unable to initialize logger", err)
	}
	ctx := context.Background()
	repository := inmemory.NewStorage()
	packagingService := service.NewService(repository)
	router := handler.Handler(packagingService)
	httpServer := server.NewHttpServer(router)

	go func() {
		logging.Logger.WithContext(ctx).Info("starting server")
		if err := httpServer.Run(); err != nil {
			logging.Logger.WithContext(ctx).Errorf("unable to start server", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c

	logging.Logger.WithContext(ctx).Infof("received signal %s", sig)
	ctx, cancel := context.WithTimeout(ctx, config.WaitingTimeout)
	defer cancel()

	err = httpServer.Stop(ctx)
	if err != nil {
		logging.Logger.WithContext(ctx).Errorf("unable to stop server gracefully", err)
	}
}
