package farmer

import (
	"chicken-farmer/backend/internal/farm/ctxfarm"
	"chicken-farmer/backend/internal/pkg"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"chicken-farmer/backend/internal/pkg/jwt"
	"context"

	"github.com/google/uuid"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authIgnoredMethods = []string{
		"/chicken_farmer.v1.FarmerPublicService/Login",
		"/chicken_farmer.v1.FarmerPublicService/Register",
	}
)

type IController interface {
	Register(
		ctx context.Context, farmerName, farmName, password string,
	) (Farmer, error)
	Login(
		ctx context.Context, farmerName, password string,
	) (jwt string, err error)
	GetGoldEggs(ctx context.Context, farmerID uuid.UUID) (uint, error)
	SpendGoldEggs(ctx context.Context, farmerID uuid.UUID, amount uint) error
	GrantGoldEggs(ctx context.Context, farmerID uuid.UUID, amount uint) error
}

type GRPCService struct {
	internalGrpc.UnimplementedFarmerPrivateServiceServer
	internalGrpc.UnimplementedFarmerPublicServiceServer

	address string
	server  *grpc.Server
	logger  *zap.SugaredLogger

	jwtAuthKey []byte

	controller IController
}

var (
	_ internalGrpc.FarmerPrivateServiceServer = &GRPCService{}
	_ internalGrpc.FarmerPublicServiceServer  = &GRPCService{}
)

func ProvideGRPCService(
	address string,
	logger *zap.SugaredLogger,
	controller IController,
	jwtAuthKey []byte,
) *GRPCService {
	return &GRPCService{
		address:    address,
		logger:     logger,
		controller: controller,
		jwtAuthKey: jwtAuthKey,
	}
}

func (s *GRPCService) ListenForConnections(ctx context.Context) {
	internalGrpc.ListenForConnections(
		ctx, s, s.address, s.logger.Desugar(), s.Authenticate,
	)
}

func (s *GRPCService) RegisterGrpcServer(server *grpc.Server) {
	// Keep track of server for the graceful stop.
	s.server = server

	internalGrpc.RegisterFarmerPrivateServiceServer(server, s)
	internalGrpc.RegisterFarmerPublicServiceServer(server, s)
}

func (s *GRPCService) GracefulStop() {
	s.logger.Info("Stopping gracefully...")
	s.server.GracefulStop()
	s.logger.Info("Stopped")
}

// Authenticate is an implementation of grpcAuth.AuthFunc specific for this
// service.
func (s *GRPCService) Authenticate(ctx context.Context) (context.Context, error) {
	// Skip authentication for specified methods.
	method, _ := grpc.Method(ctx)
	if slices.Index(authIgnoredMethods, method) != -1 {
		return ctx, nil
	}

	bearerToken, err := grpcAuth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		s.logger.Debug(err)
		return nil, internalGrpc.ErrMissingMetadata
	}

	claims, err := jwt.ValidateUserClaims(s.jwtAuthKey, bearerToken)
	if err != nil {
		s.logger.Debug(err)
		return nil, internalGrpc.ErrInvalidToken
	}

	// TODO ctxfarmer ?
	return ctxfarm.SetInContext(ctx, claims.FarmerID, claims.FarmID), nil
}

func (s *GRPCService) Register(
	ctx context.Context, request *internalGrpc.RegisterRequest,
) (*internalGrpc.RegisterResponse, error) {
	farmer, err := s.controller.Register(
		ctx,
		request.GetFarmerName(),
		request.GetFarmName(),
		request.GetPassword(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.RegisterResponse{
		FarmerId: farmer.ID.String(),
		FarmId:   farmer.FarmID.String(),
	}, nil
}

func (s *GRPCService) Login(
	ctx context.Context, request *internalGrpc.LoginRequest,
) (*internalGrpc.LoginResponse, error) {
	authToken, err := s.controller.Login(
		ctx, request.GetFarmerName(), request.GetPassword(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.LoginResponse{
		AuthToken: authToken,
	}, nil
}

func (s *GRPCService) GrantGoldEggs(
	ctx context.Context, request *internalGrpc.GrantGoldEggsRequest,
) (*internalGrpc.GrantGoldEggsResponse, error) {
	if err := s.controller.GrantGoldEggs(
		ctx,
		pkg.UUIDFromString(request.GetFarmerId()),
		uint(request.GetAmount()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.GrantGoldEggsResponse{}, nil
}

func (s *GRPCService) SpendGoldEggs(
	ctx context.Context, request *internalGrpc.SpendGoldEggsRequest,
) (*internalGrpc.SpendGoldEggsResponse, error) {
	if err := s.controller.SpendGoldEggs(
		ctx,
		pkg.UUIDFromString(request.GetFarmerId()),
		uint(request.GetAmount()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.SpendGoldEggsResponse{}, nil
}

func (s *GRPCService) GetGoldEggs(
	ctx context.Context, request *internalGrpc.GetGoldEggsRequest,
) (*internalGrpc.GetGoldEggsResponse, error) {
	goldEggCount, err := s.controller.GetGoldEggs(
		ctx, pkg.UUIDFromString(request.GetFarmerId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.GetGoldEggsResponse{Amount: uint32(goldEggCount)}, nil
}
