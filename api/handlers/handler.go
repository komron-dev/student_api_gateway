package handler

import (
	"student_api_gateway/config"
	"student_api_gateway/pkg/logger"
	"student_api_gateway/services"
)

type HandlerConfig struct {
	Logger            logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
}

func New(c *HandlerConfig) *HandlerConfig {
	return &HandlerConfig{
		Logger: c.Logger,
		ServiceManager: c.ServiceManager,
		Cfg: c.Cfg,
	}
}
