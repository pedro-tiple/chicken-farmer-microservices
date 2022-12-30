package grpc

import (
	"context"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceRegistrar interface {
	RegisterGrpcServer(server *grpc.Server)
}

// ListenForConnections starts a gRPC server.
func ListenForConnections(
	ctx context.Context,
	registrar ServiceRegistrar,
	address string,
	logger *zap.Logger,
	authFunction grpc_auth.AuthFunc,
) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		slog.Error("Errored listening", err)
		return
	}

	// TODO use recovery only when in production
	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_auth.StreamServerInterceptor(authFunction),
			//grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_auth.UnaryServerInterceptor(authFunction),
			//grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	registrar.RegisterGrpcServer(srv)
	//reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		slog.Error("Errored serving", err)
		return
	}
}

// CreateClientConnection creates a gRPC client connection.
func CreateClientConnection(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, []grpc.DialOption{
		// TODO
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		slog.Error("Errored dialing", err)
		return nil
	}

	return conn
}
