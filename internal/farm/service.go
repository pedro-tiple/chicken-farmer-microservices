package farm

import (
	"chicken-farmer/backend/internal/farm/ctxfarm"
	"chicken-farmer/backend/internal/pkg"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"

	"github.com/google/uuid"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IController interface {
	GetFarm(ctx context.Context) (GetFarmResult, error)

	BuyBarn(ctx context.Context) error
	BuyFeedBag(ctx context.Context, barnID uuid.UUID, amount uint) error
	BuyChicken(ctx context.Context, barnID uuid.UUID) error

	FeedChicken(ctx context.Context, chickenID uuid.UUID) error
	FeedChickensOfBarn(ctx context.Context, barnID uuid.UUID) error

	SetDay(ctx context.Context, day uint) error
}

type Service struct {
	internalGrpc.UnimplementedFarmServiceServer

	address string
	server  *grpc.Server
	logger  *zap.SugaredLogger

	// TODO message queue to receive time ticks.
	controller IController
}

var _ internalGrpc.FarmServiceServer = &Service{}

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
	// 	return nil, err
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

	internalGrpc.RegisterFarmServiceServer(server, s)
}

func (s *Service) GracefulStop() {
	s.logger.Info("Stopping gracefully...")
	s.server.GracefulStop()
	s.logger.Info("Stopped")
}

func (s *Service) GetFarm(
	ctx context.Context, _ *internalGrpc.GetFarmRequest,
) (*internalGrpc.GetFarmResponse, error) {
	farm, err := s.controller.GetFarm(ctx)
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

	return &internalGrpc.GetFarmResponse{
		Farm: &internalGrpc.Farm{
			Name:       farm.Name,
			Barns:      protoBarns,
			Day:        uint32(farm.CurrentDay),
			GoldenEggs: uint32(farm.GoldEggCount),
		},
	}, nil
}

func (s *Service) BuyBarn(
	ctx context.Context, _ *internalGrpc.BuyBarnRequest,
) (*internalGrpc.BuyBarnResponse, error) {
	if err := s.controller.BuyBarn(ctx); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.BuyBarnResponse{}, nil
}

func (s *Service) BuyFeedBag(
	ctx context.Context, request *internalGrpc.BuyFeedBagRequest,
) (*internalGrpc.BuyFeedBagResponse, error) {
	if err := s.controller.BuyFeedBag(ctx,
		pkg.UUIDFromString(request.GetBarnId()),
		uint(request.GetAmount()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.BuyFeedBagResponse{}, nil
}

func (s *Service) BuyChicken(
	ctx context.Context, request *internalGrpc.BuyChickenRequest,
) (*internalGrpc.BuyChickenResponse, error) {
	if err := s.controller.BuyChicken(
		ctx, pkg.UUIDFromString(request.GetBarnId()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.BuyChickenResponse{}, nil
}

func (s *Service) FeedChicken(
	ctx context.Context, request *internalGrpc.FeedChickenRequest,
) (*internalGrpc.FeedChickenResponse, error) {
	if err := s.controller.FeedChicken(
		ctx, pkg.UUIDFromString(request.GetChickenId()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.FeedChickenResponse{}, nil
}

func (s *Service) FeedChickensOfBarn(
	ctx context.Context, request *internalGrpc.FeedChickensOfBarnRequest,
) (*internalGrpc.FeedChickensOfBarnResponse, error) {
	if err := s.controller.FeedChickensOfBarn(
		ctx, pkg.UUIDFromString(request.GetBarnId()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &internalGrpc.FeedChickensOfBarnResponse{}, nil
}
