package main

import (
	// "fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/jaloldinov/Udevs_task/api_gateway/api"
	"github.com/jaloldinov/Udevs_task/api_gateway/config"
	"github.com/jaloldinov/Udevs_task/api_gateway/pkg/logger"
	"github.com/jaloldinov/Udevs_task/api_gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	gprcClients, _ := services.NewGrpcClients(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: gprcClients,
	})

	quit := make(chan os.Signal, 1)
	go server.Run(cfg.HttpPort)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	log.Info("Server exiting")
}
