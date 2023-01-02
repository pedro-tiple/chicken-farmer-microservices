package farm

import (
	"chicken-farmer/backend/internal/farm/ctxfarm"
	farmGrpc "chicken-farmer/backend/internal/farm/grpc"
	"chicken-farmer/backend/internal/pkg"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
	"fmt"

	"github.com/google/uuid"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

var _ farmGrpc.FarmServiceServer = &Service{}

type Service struct {
	farmGrpc.UnimplementedFarmServiceServer

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

	farmGrpc.RegisterFarmServiceServer(server, s)
}

func (s *Service) GracefulStop() {
	slog.Info("Stopping gracefully...")
	s.server.GracefulStop()
	slog.Info("Stopped")
}

func (s *Service) GetFarm(
	ctx context.Context, _ *farmGrpc.GetFarmRequest,
) (*farmGrpc.GetFarmResponse, error) {
	farm, err := s.controller.GetFarm(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	protoBarns := make([]*farmGrpc.Barn, len(farm.Barns))
	for i, barn := range farm.Barns {
		protoChickens := make([]*farmGrpc.Chicken, len(barn.Chickens))
		for t, chicken := range barn.Chickens {
			protoChickens[t] = &farmGrpc.Chicken{
				Id:             chicken.ID.String(),
				DateOfBirth:    uint32(chicken.DateOfBirth),
				RestingUntil:   uint32(chicken.RestingUntil),
				NormalEggsLaid: uint32(chicken.NormalEggsLaid),
				GoldEggsLaid:   uint32(chicken.GoldEggsLaid),
			}
		}

		protoBarns[i] = &farmGrpc.Barn{
			Id:            barn.ID.String(),
			Feed:          uint32(barn.Feed),
			HasAutoFeeder: barn.HasAutoFeeder,
			Chickens:      protoChickens,
		}
	}

	return &farmGrpc.GetFarmResponse{
		Farm: &farmGrpc.Farm{
			Name:       farm.Name,
			Barns:      protoBarns,
			Day:        uint32(farm.CurrentDay),
			GoldenEggs: uint32(farm.GoldEggCount),
		},
	}, nil
}

func (s *Service) BuyBarn(
	ctx context.Context, _ *farmGrpc.BuyBarnRequest,
) (*farmGrpc.BuyBarnResponse, error) {
	return &farmGrpc.BuyBarnResponse{}, s.controller.BuyBarn(ctx)
}

func (s *Service) BuyFeed(
	ctx context.Context, request *farmGrpc.BuyFeedRequest,
) (*farmGrpc.BuyFeedResponse, error) {
	fmt.Println("buyfeed", request.GetBarnId(), request.GetAmount())
	return &farmGrpc.BuyFeedResponse{}, s.controller.BuyFeed(
		ctx,
		pkg.UUIDFromString(request.GetBarnId()),
		uint(request.GetAmount()),
	)
}

func (s *Service) BuyChicken(
	ctx context.Context, request *farmGrpc.BuyChickenRequest,
) (*farmGrpc.BuyChickenResponse, error) {
	return &farmGrpc.BuyChickenResponse{}, s.controller.BuyChicken(
		ctx, pkg.UUIDFromString(request.GetBarnId()),
	)
}

func (s *Service) FeedChicken(
	ctx context.Context, request *farmGrpc.FeedChickenRequest,
) (*farmGrpc.FeedChickenResponse, error) {
	return &farmGrpc.FeedChickenResponse{}, s.controller.FeedChicken(
		ctx, pkg.UUIDFromString(request.GetChickenId()),
	)
}

func (s *Service) FeedChickensOfBarn(
	ctx context.Context, request *farmGrpc.FeedChickensOfBarnRequest,
) (*farmGrpc.FeedChickensOfBarnResponse, error) {
	return &farmGrpc.FeedChickensOfBarnResponse{}, s.controller.FeedChickensOfBarn(
		ctx, pkg.UUIDFromString(request.GetBarnId()),
	)
}
