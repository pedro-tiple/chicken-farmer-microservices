package farm

import (
	"chicken-farmer/backend/internal/farm/ctxfarm"
	"chicken-farmer/backend/internal/pkg"
	cfGrpc "chicken-farmer/backend/internal/pkg/grpc"
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
	cfGrpc.UnimplementedFarmServiceServer

	address string
	server  *grpc.Server
	logger  *zap.SugaredLogger

	// TODO message queue to receive time ticks.
	controller IController
}

var _ cfGrpc.FarmServiceServer = &Service{}

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
	//token, err := grpcAuth.AuthFromMD(ctx, "bearer")
	//if err != nil {
	//	return nil, err
	//}

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

	cfGrpc.RegisterFarmServiceServer(server, s)
}

func (s *Service) GracefulStop() {
	s.logger.Info("Stopping gracefully...")
	s.server.GracefulStop()
	s.logger.Info("Stopped")
}

func (s *Service) GetFarm(
	ctx context.Context, _ *cfGrpc.GetFarmRequest,
) (*cfGrpc.GetFarmResponse, error) {
	farm, err := s.controller.GetFarm(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	protoBarns := make([]*cfGrpc.Barn, len(farm.Barns))
	for i, barn := range farm.Barns {
		protoChickens := make([]*cfGrpc.Chicken, len(barn.Chickens))
		for t, chicken := range barn.Chickens {
			protoChickens[t] = &cfGrpc.Chicken{
				Id:             chicken.ID.String(),
				DateOfBirth:    uint32(chicken.DateOfBirth),
				RestingUntil:   uint32(chicken.RestingUntil),
				NormalEggsLaid: uint32(chicken.NormalEggsLaid),
				GoldEggsLaid:   uint32(chicken.GoldEggsLaid),
			}
		}

		protoBarns[i] = &cfGrpc.Barn{
			Id:            barn.ID.String(),
			Feed:          uint32(barn.Feed),
			HasAutoFeeder: barn.HasAutoFeeder,
			Chickens:      protoChickens,
		}
	}

	return &cfGrpc.GetFarmResponse{
		Farm: &cfGrpc.Farm{
			Name:       farm.Name,
			Barns:      protoBarns,
			Day:        uint32(farm.CurrentDay),
			GoldenEggs: uint32(farm.GoldEggCount),
		},
	}, nil
}

func (s *Service) BuyBarn(
	ctx context.Context, _ *cfGrpc.BuyBarnRequest,
) (*cfGrpc.BuyBarnResponse, error) {
	if err := s.controller.BuyBarn(ctx); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &cfGrpc.BuyBarnResponse{}, nil
}

func (s *Service) BuyFeedBag(
	ctx context.Context, request *cfGrpc.BuyFeedBagRequest,
) (*cfGrpc.BuyFeedBagResponse, error) {
	if err := s.controller.BuyFeedBag(ctx,
		pkg.UUIDFromString(request.GetBarnId()),
		uint(request.GetAmount()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &cfGrpc.BuyFeedBagResponse{}, nil
}

func (s *Service) BuyChicken(
	ctx context.Context, request *cfGrpc.BuyChickenRequest,
) (*cfGrpc.BuyChickenResponse, error) {
	if err := s.controller.BuyChicken(
		ctx, pkg.UUIDFromString(request.GetBarnId()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &cfGrpc.BuyChickenResponse{}, nil
}

func (s *Service) FeedChicken(
	ctx context.Context, request *cfGrpc.FeedChickenRequest,
) (*cfGrpc.FeedChickenResponse, error) {
	if err := s.controller.FeedChicken(
		ctx, pkg.UUIDFromString(request.GetChickenId()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &cfGrpc.FeedChickenResponse{}, nil
}

func (s *Service) FeedChickensOfBarn(
	ctx context.Context, request *cfGrpc.FeedChickensOfBarnRequest,
) (*cfGrpc.FeedChickensOfBarnResponse, error) {
	if err := s.controller.FeedChickensOfBarn(
		ctx, pkg.UUIDFromString(request.GetBarnId()),
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &cfGrpc.FeedChickensOfBarnResponse{}, nil
}
