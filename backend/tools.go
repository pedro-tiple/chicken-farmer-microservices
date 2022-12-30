//go:build tools
// +build tools

// following https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package backend

import (
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
)
