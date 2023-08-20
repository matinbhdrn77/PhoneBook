package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZap(cfg *Config) *zap.Logger {
	return zap.New(
		zapcore.NewCore(getEncoder(cfg), getWriteSyncer(cfg), getLoggerLevel(cfg)),
		getOptions(cfg)...,
	)
}
