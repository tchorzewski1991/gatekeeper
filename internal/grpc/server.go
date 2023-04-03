package grpc

import (
	"context"
	"github.com/tchorzewski1991/gatekeeper/pkg/gk"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	gk.UnimplementedApiServer
	grpcSrv *grpc.Server
}

func NewServer() *Server {
	var s Server

	grpcSrv := grpc.NewServer()

	gk.RegisterApiServer(grpcSrv, &s)

	s.grpcSrv = grpcSrv

	return &s
}

func Run(lc fx.Lifecycle, server *Server, logger *zap.Logger) error {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			l, err := net.Listen("tcp", ":8081")
			if err != nil {
				return err
			}

			logger.Info("Starting grpc server", zap.String("addr", "8081"))

			go func() {
				if err = server.grpcSrv.Serve(l); err != nil {
					logger.Error("Unable to start grpc server", zap.String("addr", "8081"))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping grpc server", zap.String("addr", ":8081"))
			server.grpcSrv.GracefulStop()
			return nil
		},
	})

	return nil
}

func (s *Server) Authenticate(_ context.Context, _ *gk.AuthRequest) (*gk.AuthResponse, error) {
	return &gk.AuthResponse{
		Token:     "token",
		ExpiresIn: 100,
	}, nil
}
