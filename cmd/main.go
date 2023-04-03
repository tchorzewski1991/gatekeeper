package main

import (
	"github.com/tchorzewski1991/gatekeeper/internal/grpc"
	"github.com/tchorzewski1991/gatekeeper/internal/log"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Provide(
			grpc.NewServer,
			log.NewLogger,
		),
		fx.Invoke(grpc.Run),
	).Run()
}
