include .env

export APP_PORT
export APP_DSN
export APP_SERVICE_NAME
export APP_LOG_LEVEL
export APP_SALT
export APP_TOKEN_TTL
export APP_REFRESH_TOKEN_TTL
export APP_TOKEN_SECRET
export APP_REDIS_URL
export APP_CACHE_DURATION

up:
	@docker-compose up -d & disown
up-dev:
	@docker-compose -f docker-compose-dev.yml up -d & disown
run:
	@cd backend && go run cmd/app/main.go
#@swagger generate server -f ./api/openapi/openapi.yaml --server-package=./internal/generated -A "backend-service"
generate:
	@rm -rf admin/src/api && openapi-generator-cli generate -i ./backend/api/openapi/openapi.yaml -o admin/src/api -g typescript-axios --additional-properties=supportsES6=true,npmVersion=6.9.0,typescriptThreePlus=true
test:
	@cd backend && go clean -testcache
	@cd backend && go test ./... -v