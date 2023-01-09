#!/bin/sh

mockgen -package farm -source ./internal/farm/controller.go -destination ./internal/farm/controller_mock.go
mockgen -package farm -source ./internal/farm/service_grpc.go -destination ./internal/farm/service_grpc_mock.go
