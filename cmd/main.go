package main

import (
	"student_api_gateway/api"
	"student_api_gateway/config"
	"student_api_gateway/pkg/logger"
	"student_api_gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "student-api-gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	err = server.Run(cfg.HTTPPort)
	if err != nil {
		log.Fatal("failed to run HTTP server", logger.Error(err))
		panic(err)
	}
}