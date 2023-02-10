package grpc

import (
	"context"
	"net"
	"net/http"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const readHeaderTimeout = 3 * time.Second

var (
	ErrMissingMetadata = status.Errorf(
		codes.Unauthenticated, "missing metadata",
	)
	ErrInvalidToken = status.Errorf(codes.PermissionDenied, "invalid token")
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
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_opentracing.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				grpc_zap.StreamServerInterceptor(logger),
				grpc_auth.StreamServerInterceptor(authFunction),
				// grpc_recovery.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_auth.UnaryServerInterceptor(authFunction),
				// grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)
	registrar.RegisterGrpcServer(srv)
	// reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		slog.Error("Errored serving", err)

		return
	}
}

// CreateClientConnection creates a gRPC client connection.
func CreateClientConnection(
	ctx context.Context, address string,
) (*grpc.ClientConn, error) {
	// TODO set secure credentials
	return grpc.DialContext(
		ctx, address, []grpc.DialOption{
			// grpc.WithReturnConnectionError(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}...,
	)
}

func RunRESTGateway(
	ctx context.Context,
	logger *zap.SugaredLogger,
	handler func(
		context.Context, *runtime.ServeMux, string, []grpc.DialOption,
	) error,
	restAddr, grpcAddr string,
) {
	mux := runtime.NewServeMux()

	if err := handler(
		ctx, mux, grpcAddr, []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		logger.Fatal(err)
	}

	server := &http.Server{
		Addr: restAddr,
		Handler: cors.New(
			cors.Options{
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut},
				AllowedHeaders:   []string{"Origin", "Authorization", "Content-Type"},
				AllowCredentials: false,
				Debug:            true,
			},
		).Handler(mux),
		ReadHeaderTimeout: readHeaderTimeout,
	}
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
