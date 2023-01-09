all: sqlc buf mockgen generate apiclient
backend: sqlc buf mockgen generate
frontend: apiclient

sqlc:
	sqlc generate -f "./internal/farm/sql/sqlc.yaml"

buf:
	cd ./api/proto; buf lint
	cd ./api/proto; buf generate

generate:
	go generate "./..."

mockgen:
	./scripts/mockgen.sh

apiclient:
	docker run --rm \
		-v "${CURDIR}:/local" \
		openapitools/openapi-generator-cli generate \
        -i /local/api/openapiv2/chicken_farmer/v1/chicken_farmer.swagger.json \
        -g typescript-axios \
        -o "$/local/web/chicken-farmer-service" \
        -p npmName=chicken-farmer-service
	cd "./web/chicken-farmer-service"; npm run build


.PHONY: sqlc buf generate mockgen apiclient
