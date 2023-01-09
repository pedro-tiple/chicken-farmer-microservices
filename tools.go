//go:build tools
// +build tools

// following https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package backend

import (
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
	_ "golang.org/x/tools/cmd/goimports"
)
