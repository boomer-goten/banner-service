package main

import (
	"banner-server/app"
	"banner-server/internal/api"
	"banner-server/internal/env"
	"banner-server/internal/logger"
	cacheinmemory "banner-server/pkg/cache/cacheInMemory"
	"banner-server/pkg/repository/postgres"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	env.LoadEnv()
	logger := logger.SetupLogger()
	cache := cacheinmemory.New()
	database := postgres.New()
	DefaultAPIService := api.NewDefaultAPIService(database, cache)
	DefaultAPIController := api.NewDefaultAPIController(DefaultAPIService)

	router := api.NewRouter(logger, DefaultAPIController)
	server := app.New(logger, router)
	go server.Run()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	server.Stop()
}
