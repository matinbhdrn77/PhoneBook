package config

import "github.com/matinbhdrn77/PhoneBook/pkg/logger"

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
