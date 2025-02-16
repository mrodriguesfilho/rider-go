package logger

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLoggerAdapter struct {
	logger *zap.SugaredLogger
}

func newZapLoggerAdapter(lc fx.Lifecycle) (CustomLogger, error) {
	zapLoggerConfig := zap.NewProductionConfig()
	zapLoggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger, err := zapLoggerConfig.Build()

	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return zapLogger.Sync()
		},
	})

	return &zapLoggerAdapter{logger: zapLogger.Sugar()}, nil
}

func (z *zapLoggerAdapter) Info(msg string, fields ...interface{}) {
	z.logger.Infow(msg, fields...)
}

func (z *zapLoggerAdapter) Error(msg string, err error, fields ...interface{}) {
	z.logger.Errorw(msg, fields...)
}
