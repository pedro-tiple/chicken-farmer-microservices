package farm

import (
	"chicken-farmer/backend/farm/ctxFarm"
	"chicken-farmer/backend/farm/proto"
	"chicken-farmer/backend/internal"
	internalGrpc "chicken-farmer/backend/internal/grpc"
	"context"

	"github.com/google/uuid"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type IController interface {
	GetFarm(ctx context.Context) (GetFarmResult, error)

	BuyBarn(ctx context.Context) error
	BuyFeed(ctx context.Context, barnID uuid.UUID, amount uint) error
	BuyChicken(ctx context.Context, barnID uuid.UUID) error

	FeedChicken(ctx context.Context, chickenID uuid.UUID) error
	FeedChickensOfBarn(ctx context.Context, barnID uuid.UUID) error

	SetDay(ctx context.Context, day uint) error
}

type Service struct {
	proto.UnimplementedFarmServiceServer

	address string
	server  *grpc.Server
	logger  *zap.SugaredLogger

	// TODO message queue to receive time ticks.
	controller IController
}

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

// Authenticate is an implementation of grpc_auth.AuthFunc specific for this
// service. We'll need one per service because of the different context values
// needed, maybe.
func Authenticate(ctx context.Context) (context.Context, error) {
	//token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	//if err != nil {
	//	return nil, err
	//}

	// TODO validate JWT and build context from claims.

	return ctxFarm.SetInContext(
		ctx,
		internal.UUIDFromString("65e4d8ff-8766-48a7-bfcd-7160d149a319"),
		internal.UUIDFromString("93020a42-c32a-4b2c-a4b9-779f82841b11"),
	), nil
}

func (s *Service) ListenForConnections(
	ctx context.Context, authFunction grpc_auth.AuthFunc,
) {
	internalGrpc.ListenForConnections(
		ctx, s, s.address, s.logger.Desugar(), authFunction,
	)
}

func (s *Service) RegisterGrpcServer(server *grpc.Server) {
	// Keep track of server for the graceful stop.
	s.server = server

	proto.RegisterFarmServiceServer(server, s)
}

func (s *Service) GracefulStop() {
	slog.Info("Stopping gracefully...")
	s.server.GracefulStop()
	slog.Info("Stopped")
}

func (s *Service) GetFarm(
	ctx context.Context, _ *proto.GetFarmRequest,
) (*proto.GetFarmResponse, error) {
	farm, err := s.controller.GetFarm(ctx)
	if err != nil {
		return nil, err
	}

	protoBarns := make([]*proto.Barn, len(farm.Barns))
	for i, barn := range farm.Barns {
		protoChickens := make([]*proto.Chicken, len(barn.Chickens))
		for t, chicken := range barn.Chickens {
			protoChickens[t] = &proto.Chicken{
				Id:             chicken.ID.String(),
				DateOfBirth:    uint32(chicken.DateOfBirth),
				RestingUntil:   uint32(chicken.RestingUntil),
				NormalEggsLaid: uint32(chicken.NormalEggsLaid),
				GoldEggsLaid:   uint32(chicken.GoldEggsLaid),
			}
		}

		protoBarns[i] = &proto.Barn{
			Id:            barn.ID.String(),
			Feed:          uint32(barn.Feed),
			HasAutoFeeder: barn.HasAutoFeeder,
			Chickens:      protoChickens,
		}
	}

	return &proto.GetFarmResponse{
		Farm: &proto.Farm{
			Name:       farm.Name,
			Barns:      protoBarns,
			Day:        uint32(farm.CurrentDay),
			GoldenEggs: uint32(farm.GoldEggCount),
		},
	}, nil
}

func (s *Service) BuyBarn(ctx context.Context, _ *proto.BuyBarnRequest) (*proto.BuyBarnResponse, error) {
	return &proto.BuyBarnResponse{}, s.controller.BuyBarn(ctx)
}

func (s *Service) BuyFeed(ctx context.Context, request *proto.BuyFeedRequest) (*proto.BuyFeedResponse, error) {
	return &proto.BuyFeedResponse{}, s.controller.BuyFeed(
		ctx,
		internal.UUIDFromString(request.GetBarnID()),
		uint(request.GetAmount()),
	)
}

func (s *Service) BuyChicken(
	ctx context.Context, request *proto.BuyChickenRequest,
) (*proto.BuyChickenResponse, error) {
	return &proto.BuyChickenResponse{}, s.controller.BuyChicken(
		ctx, internal.UUIDFromString(request.GetBarnID()),
	)
}

func (s *Service) FeedChicken(
	ctx context.Context, request *proto.FeedChickenRequest,
) (*proto.FeedChickenResponse, error) {
	return &proto.FeedChickenResponse{}, s.controller.FeedChicken(
		ctx, internal.UUIDFromString(request.GetChickenID()),
	)
}

func (s *Service) FeedChickensOfBarn(
	ctx context.Context, request *proto.FeedChickensOfBarnRequest,
) (*proto.FeedChickensOfBarnResponse, error) {
	return &proto.FeedChickensOfBarnResponse{}, s.controller.FeedChickensOfBarn(
		ctx, internal.UUIDFromString(request.GetBarnID()),
	)
}
