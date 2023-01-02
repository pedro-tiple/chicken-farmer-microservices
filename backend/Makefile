all: sqlc buf mockgen generate

sqlc:
	sqlc generate -f ./internal/farm/sql/sqlc.yaml

buf:
	cd api/proto; buf generate

generate:
	go generate ./...

mockgen:
	"$(CURDIR)/scripts/mockgen.sh"

.PHONY: sqlc buf generate mockgen