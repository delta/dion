// Package to handle all the logging
// TODO
// 1. Production configuration
// 2. Log file splitting
// 3. Tests
// 4. docs

package logging

import (
	"os"

	"delta.nitt.edu/dion/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zap.Logger instance which can be used to create logs
//
// L is only instantiated only when the Setup() function is called
var L *zap.Logger

// parameters which needs to be passed to Setup function when
// logger is being created
type loggerOption struct {
	Level      zapcore.LevelEnabler
	zapOptions []zap.Option
}

// Interface for all the logging option to be passed to
// comply with zap.LoggingOptions
type LoggerOptions interface {
	applyLoggerOption(l *loggerOption)
}

type loggerOptionFunc func(*loggerOption)

func (f loggerOptionFunc) applyLoggerOption(opts *loggerOption) {
	f(opts)
}

// WrapOptions adds zap.Option's to a test Logger built by NewLogger.
func WrapOptions(zapOpts ...zap.Option) LoggerOptions {
	return loggerOptionFunc(func(opts *loggerOption) {
		opts.zapOptions = zapOpts
	})
}

func Level(level zapcore.LevelEnabler) LoggerOptions {
	return loggerOptionFunc(func(opts *loggerOption) {
		opts.Level = level
	})
}

// Initializes the logger according to the environment
func Setup(opts ...LoggerOptions) {
	if config.C.Environment == "dev" {
		SetupDevLogger(opts...)
	} else {
		SetupProdLogger(opts...)
	}
}

func SetupDevLogger(opts ...LoggerOptions) {
	cfg := &loggerOption{
		Level: zapcore.DebugLevel,
	}

	for _, o := range opts {
		o.applyLoggerOption(cfg)
	}

	// writer := zapcore.AddSync(io.Discard)

	L = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(
				zap.NewDevelopmentEncoderConfig(),
			),
			os.Stdout,
			cfg.Level,
		),
		cfg.zapOptions...,
	)
}

func SetupProdLogger(opts ...LoggerOptions) {
	// setup a prod logger with given data
	cfg := &loggerOption{
		Level: zapcore.InfoLevel,
	}

	for _, o := range opts {
		o.applyLoggerOption(cfg)
	}

	L = zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(
				zap.NewProductionEncoderConfig(),
			),
			os.Stdout,
			cfg.Level,
		),
		cfg.zapOptions...,
	)
}

// Flushes the logger's buffer
//
// It should be called after Setup as a defer,
//
//	logging.Setup(...)
//	defer logging.Flush()
func Flush() {
	L.Sync()
}

// Method to access sugared logger
//
// Instead of
//
//	logging.L.Sugar().Info("hello world")
//
// use,
//
//	logging.Sugared().Info("hello world")
//
// this makes the code cleaner
func Sugared() *zap.SugaredLogger {
	return L.Sugar()
}
