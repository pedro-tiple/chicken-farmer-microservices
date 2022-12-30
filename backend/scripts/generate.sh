# This file will generate all mocks, proto files.

# Sqlc
sqlc generate -f ./internal/farm/sql/sqlc.yaml

# Protobufs
protoc --go_opt=module=github.com/pedro-tiple/proto-farmer-ethereum/backend/internal/farm/grpc --go-grpc_opt=module=github.com/pedro-tiple/proto-farmer-ethereum/backend/internal/farm/grpc --go_out=./internal/farm/grpc --go-grpc_out=./internal/farm/grpc ./api/proto/farm.proto

# Mocks
mockgen -package farm -source ./internal/farm/controller.go -destination ./internal/farm/controller_mock.go
mockgen -package farm -source ./internal/farm/service.go -destination ./internal/farm/service_mock.go

go generate ./...
