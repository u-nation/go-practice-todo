package main

import (
	"github.com/u-nation/go-practice-todo/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/xerrors"
)

func setupLogger(conf *config.APIConfig) error {
	var level = zapcore.DebugLevel

	if conf.IsPrd() {
		level = zapcore.InfoLevel
	}

	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(level),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "severity",
			TimeKey:      "time",
			MessageKey:   "message",
			CallerKey:    "file",
			EncodeTime:   zapcore.RFC3339NanoTimeEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	newLogger, err := logConfig.Build()
	if err != nil {
		return xerrors.Errorf("logConfig.Build: %w", err)
	}
	zap.ReplaceGlobals(newLogger)

	zap.L().Info("Logger setting completed")
	return nil
}
