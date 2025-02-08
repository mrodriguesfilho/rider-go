package logger

import (
	"go.uber.org/fx"
)

type CustomLogger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, err error, fields ...interface{})
}

func NewLogger(lc fx.Lifecycle) CustomLogger {

	logger, err := newZapLoggerAdapter(lc)

	if err != nil {
		panic(err)
	}

	return logger
}
