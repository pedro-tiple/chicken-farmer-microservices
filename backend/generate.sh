# This file will generate all mocks, proto files.

# Sqlc
sqlc generate -f ./farm/sql/sqlc.yaml

# Protobufs
protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=. --go-grpc_out=. ./farm/proto/farm.proto

# Mocks
mockgen -package farm -source ./farm/controller.go -destination ./farm/controller_mock.go
mockgen -package farm -source ./farm/service.go -destination ./farm/service_mock.go

go generate ./...
