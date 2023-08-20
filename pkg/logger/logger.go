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

func getEncoder(cfg *Config) zapcore.Encoder {
	var encodeConfig zapcore.EncoderConfig
	if cfg.Development {
		encodeConfig = zap.NewDevelopmentEncoderConfig()
		encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encodeConfig = zap.NewProductionEncoderConfig()
		encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encodeConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encodeConfig)
	}

	return encoder
}
