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
run:
	@go run cmd/app/main.go
swaggen:
	@swagger generate server -f ./api/openapi/openapi.yml --server-package=./internal/generated -A "backend-service"
test:
	@go clean -testcache
	@go test ./... -v