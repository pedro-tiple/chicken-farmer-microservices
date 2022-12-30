package grpc

import (
	"chicken-farmer/backend/internal/pkg/grpc"
)

func NewClient(address string) FarmServiceClient {
	return NewFarmServiceClient(grpc.CreateClientConnection(address))
}
