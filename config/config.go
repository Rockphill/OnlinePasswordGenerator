package config

import (
	"os"
	"time"
)

type Config struct {
	Port          string
	Environment   string
	StartTime     time.Time
	TemplatesPath string
}

func NewConfig() *Config {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	return &Config{
		Port:          os.Getenv("PORT"),
		Environment:   env,
		StartTime:     time.Now(),
		TemplatesPath: "templates/",
	}
}
