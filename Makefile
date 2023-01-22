all: sqlc buf mockgen gogen
backend: sqlc buf mockgen generate
frontend: buf

sqlc:
	sqlc generate -f "./internal/farm/sql/sqlc.yaml"

buf:
	cd ./api/proto; buf lint
	cd ./api/proto; buf generate
# Proto to typescript-axios client.
	docker run --rm \
		-v "${CURDIR}:/local" \
		openapitools/openapi-generator-cli generate \
        -i /local/api/openapiv2/chicken_farmer/v1/chicken_farmer.swagger.json \
        -g typescript-axios \
        -o "$/local/web/chicken-farmer-service" \
        -p npmName=chicken-farmer-service
# Regenerate dist on.
	cd "./web/chicken-farmer-service"; npm run build
# Clear node_modules import and install again.
	cd "./web/react-ts"; rm -rf node_modules/chicken-farmer-service node_modules/.vite; npm install

gogen:
	go generate "./..."

mockgen:
	./scripts/mockgen.sh

.PHONY: sqlc buf gogen mockgen
