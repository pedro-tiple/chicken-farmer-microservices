package farm

import (
	"chicken-farmer/backend/internal/farm/ctxfarm"
	"chicken-farmer/backend/internal/pkg"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"chicken-farmer/backend/internal/pkg/jwt"
	"context"
	"fmt"

	"github.com/google/uuid"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authIgnoredMethods = []string{}
)

type IController interface {
	NewFarm(
		ctx context.Context, farmerID uuid.UUID, name string,
	) (farmID uuid.UUID, err error)
	FarmDetails(
		ctx context.Context, farmerID, farmID uuid.UUID,
	) (FarmDetailsResult, error)

	BuyBarn(ctx context.Context, farmerID, farmID uuid.UUID) error
	BuyFeedBags(
		ctx context.Context, farmerID, barnID uuid.UUID, amount uint,
	) error
	BuyChicken(ctx context.Context, farmerID, barnID uuid.UUID) error
	SellChicken(ctx context.Context, farmerID, chickenID uuid.UUID) error

	FeedChicken(ctx context.Context, farmerID, chickenID uuid.UUID) error

	SetDay(ctx context.Context, day uint) error
}

type GRPCService struct {
	internalGrpc.UnimplementedFarmServiceServer

	address string
	server  *grpc.Server
	logger  *zap.SugaredLogger

	jwtAuthKey []byte

	// TODO message queue to receive universe ticks.
	controller IController
}

var _ internalGrpc.FarmServiceServer = &GRPCService{}

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

	internalGrpc.RegisterFarmServiceServer(server, s)
}

func (s *GRPCService) GracefulStop() {
	s.logger.Info("Stopping gracefully...")
	s.server.GracefulStop()
	s.logger.Info("Stopped")
}

// Authenticate is an implementation of grpcAuth.AuthFunc specific for this
// service. We'll need one per service because of the different context values
// needed, maybe.
func (s *GRPCService) Authenticate(ctx context.Context) (context.Context, error) {
	fmt.Println("authenticate")
	// Skip authentication for specified methods.
	method, _ := grpc.Method(ctx)
	if slices.Index(authIgnoredMethods, method) != -1 {
		return ctx, nil
	}

	bearerToken, err := grpcAuth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		fmt.Println(err)
		s.logger.Info(err)
		return nil, internalGrpc.ErrMissingMetadata
	}

	fmt.Println("bearer token", bearerToken)

	claims, err := jwt.ValidateUserClaims(s.jwtAuthKey, bearerToken)
	if err != nil {
		s.logger.Info(err)
		return nil, internalGrpc.ErrInvalidToken
	}

	return ctxfarm.SetInContext(ctx, claims.FarmerID, claims.FarmID), nil
}

func (s *GRPCService) NewFarm(
	ctx context.Context, request *internalGrpc.NewFarmRequest,
) (*internalGrpc.NewFarmResponse, error) {
	// TODO validate that this is coming from a valid source.

	farmID, err := s.controller.NewFarm(
		ctx, pkg.UUIDFromString(request.GetOwnerId()), request.GetName(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.NewFarmResponse{
		FarmId: farmID.String(),
	}, nil
}

func (s *GRPCService) FarmDetails(
	ctx context.Context, _ *internalGrpc.FarmDetailsRequest,
) (*internalGrpc.FarmDetailsResponse, error) {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	farm, err := s.controller.FarmDetails(ctx, ctxData.FarmerID, ctxData.FarmID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	protoBarns := make([]*internalGrpc.Barn, len(farm.Barns))
	for i, barn := range farm.Barns {
		protoChickens := make([]*internalGrpc.Chicken, len(barn.Chickens))
		for t, chicken := range barn.Chickens {
			protoChickens[t] = &internalGrpc.Chicken{
				Id:             chicken.ID.String(),
				DateOfBirth:    uint32(chicken.DateOfBirth),
				RestingUntil:   uint32(chicken.RestingUntil),
				NormalEggsLaid: uint32(chicken.NormalEggsLaid),
				GoldEggsLaid:   uint32(chicken.GoldEggsLaid),
			}
		}

		protoBarns[i] = &internalGrpc.Barn{
			Id:            barn.ID.String(),
			Feed:          uint32(barn.Feed),
			HasAutoFeeder: barn.HasAutoFeeder,
			Chickens:      protoChickens,
		}
	}

	return &internalGrpc.FarmDetailsResponse{
		Farm: &internalGrpc.Farm{
			Name:       farm.Name,
			Barns:      protoBarns,
			Day:        uint32(farm.CurrentDay),
			GoldenEggs: uint32(farm.GoldEggCount),
		},
	}, nil
}

func (s *GRPCService) BuyBarn(
	ctx context.Context, _ *internalGrpc.BuyBarnRequest,
) (*internalGrpc.BuyBarnResponse, error) {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.controller.BuyBarn(
		ctx, ctxData.FarmerID, ctxData.FarmID,
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.BuyBarnResponse{}, nil
}

func (s *GRPCService) BuyFeedBag(
	ctx context.Context, request *internalGrpc.BuyFeedBagRequest,
) (*internalGrpc.BuyFeedBagResponse, error) {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.controller.BuyFeedBags(
		ctx,
		ctxData.FarmerID,
		pkg.UUIDFromString(request.GetBarnId()),
		uint(request.GetAmount()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.BuyFeedBagResponse{}, nil
}

func (s *GRPCService) BuyChicken(
	ctx context.Context, request *internalGrpc.BuyChickenRequest,
) (*internalGrpc.BuyChickenResponse, error) {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.controller.BuyChicken(
		ctx, ctxData.FarmerID, pkg.UUIDFromString(request.GetBarnId()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.BuyChickenResponse{}, nil
}

func (s *GRPCService) SellChicken(
	ctx context.Context, request *internalGrpc.SellChickenRequest,
) (*internalGrpc.SellChickenResponse, error) {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.controller.SellChicken(
		ctx, ctxData.FarmerID, pkg.UUIDFromString(request.GetChickenId()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.SellChickenResponse{}, nil
}

func (s *GRPCService) FeedChicken(
	ctx context.Context, request *internalGrpc.FeedChickenRequest,
) (*internalGrpc.FeedChickenResponse, error) {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.controller.FeedChicken(
		ctx, ctxData.FarmerID, pkg.UUIDFromString(request.GetChickenId()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.FeedChickenResponse{}, nil
}
