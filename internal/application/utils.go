package application

import (
	"os"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "80"
	}
	return config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}
