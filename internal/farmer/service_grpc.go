package farmer

import (
	"chicken-farmer/backend/internal/farm/ctxfarm"
	"chicken-farmer/backend/internal/pkg"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
	"fmt"

	"github.com/google/uuid"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
}

type Service struct {
	internalGrpc.UnimplementedFarmerServiceServer

	address string
	server  *grpc.Server
	logger  *zap.SugaredLogger

	controller IController
}

var _ internalGrpc.FarmerServiceServer = &Service{}

func ProvideService(
	address string,
	logger *zap.SugaredLogger,
	controller IController,
) Service {
	return Service{
		address:    address,
		logger:     logger,
		controller: controller,
	}
}

// Authenticate is an implementation of grpcAuth.AuthFunc specific for this
// service. We'll need one per service because of the different context values
// needed, maybe.
func Authenticate(ctx context.Context) (context.Context, error) {
	// token, err := grpcAuth.AuthFromMD(ctx, "bearer")
	// if err != nil {
	//	return nil, err
	// }
	// TODO validate JWT and build context from claims.
	return ctxfarm.SetInContext(
		ctx,
		pkg.UUIDFromString("65e4d8ff-8766-48a7-bfcd-7160d149a319"),
		pkg.UUIDFromString("93020a42-c32a-4b2c-a4b9-779f82841b11"),
	), nil
}

func (s *Service) ListenForConnections(
	ctx context.Context, authFunction grpcAuth.AuthFunc,
) {
	internalGrpc.ListenForConnections(
		ctx, s, s.address, s.logger.Desugar(), authFunction,
	)
}

func (s *Service) RegisterGrpcServer(server *grpc.Server) {
	// Keep track of server for the graceful stop.
	s.server = server

	internalGrpc.RegisterFarmerServiceServer(server, s)
}

func (s *Service) GracefulStop() {
	s.logger.Info("Stopping gracefully...")
	s.server.GracefulStop()
	s.logger.Info("Stopped")
}

func (s *Service) Register(
	ctx context.Context, request *internalGrpc.RegisterRequest,
) (*internalGrpc.RegisterResponse, error) {
	fmt.Println("service register")
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

func (s *Service) Login(
	ctx context.Context, request *internalGrpc.LoginRequest,
) (*internalGrpc.LoginResponse, error) {
	jwt, err := s.controller.Login(
		ctx, request.GetFarmerName(), request.GetPassword(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.LoginResponse{
		Jwt: jwt,
	}, nil
}

func (s *Service) SpendGoldEggs(
	ctx context.Context, request *internalGrpc.SpendGoldEggsRequest,
) (*internalGrpc.SpendGoldEggsResponse, error) {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.controller.SpendGoldEggs(
		ctx, ctxData.FarmerID, uint(request.GetAmount()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.SpendGoldEggsResponse{}, nil
}

func (s *Service) GetGoldEggs(
	ctx context.Context, _ *internalGrpc.GetGoldEggsRequest,
) (*internalGrpc.GetGoldEggsResponse, error) {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	goldEggCount, err := s.controller.GetGoldEggs(ctx, ctxData.FarmerID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.GetGoldEggsResponse{Amount: uint32(goldEggCount)}, nil
}
