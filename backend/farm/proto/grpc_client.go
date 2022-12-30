package proto

import "chicken-farmer/backend/internal/grpc"

func NewClient(address string) FarmServiceClient {
	return NewFarmServiceClient(grpc.CreateClientConnection(address))
}
