package main

import (
	"github.com/muhriddinsalohiddin/api-gateway/api"
	"github.com/muhriddinsalohiddin/api-gateway/config"
	"github.com/muhriddinsalohiddin/api-gateway/pkg/logger"
	"github.com/muhriddinsalohiddin/api-gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api-gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})
	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}