package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_SugarLog *zap.SugaredLogger
)

const (
	ProductionEnv  = "production"
	DevelopmentEnv = "development"
)

type LoggerConfig struct {
	Level            string
	Encoding         string
	OutputPaths      []string
	ErrorOutputPaths []string
	Env              string
}

func Init(cfg LoggerConfig) {
	var level zap.AtomicLevel
	var err error
	if level, err = zap.ParseAtomicLevel(cfg.Level); err != nil {
		panic(err)
	}

	var encoderConfig zapcore.EncoderConfig
	switch cfg.Env {
	case ProductionEnv:
		encoderConfig = zap.NewProductionEncoderConfig()
	default:
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	logger := zap.Must(zap.Config{
		Level:            level,
		Encoding:         cfg.Encoding,
		EncoderConfig:    encoderConfig,
		OutputPaths:      cfg.OutputPaths,
		ErrorOutputPaths: cfg.ErrorOutputPaths,
	}.Build())

	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	_SugarLog = logger.Sugar()
}
