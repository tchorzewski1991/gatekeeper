package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {
	conf := zap.NewProductionConfig()
	conf.DisableStacktrace = true
	conf.DisableCaller = true
	conf.OutputPaths = []string{"stdout"}
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := conf.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
