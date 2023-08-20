package config

import "github.com/matinbhdrn77/PhoneBook/pkg/logger"

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
	}
}
