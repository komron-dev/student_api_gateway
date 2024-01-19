package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Enivorentment string

	StudentServiceHost string
	StudentServicePort int

	TaskServiceHost string
	TaskServicePort int

	ContextTimeout int
	LogLevel       string
	HTTPPort       string
}

func Load() Config {
	c := Config{}

	c.Enivorentment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	
	c.StudentServiceHost = cast.ToString(getOrReturnDefault("STUDENT_SERVICE_HOST", "localhost"))
	c.StudentServicePort = cast.ToInt(getOrReturnDefault("STUDENT_SERVICE_PORT", 9005))
	
	c.TaskServiceHost = cast.ToString(getOrReturnDefault("TASK_SERVICE_HOST", "localhost"))
	c.TaskServicePort = cast.ToInt(getOrReturnDefault("TASK_SERVICE_PORT", 9006))
	
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))
	c.ContextTimeout = cast.ToInt(getOrReturnDefault("CONTEXT_TIMEOUT", 7))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
