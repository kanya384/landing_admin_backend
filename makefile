up:
	@docker-compose up -d & disown
test:
	@go clean -testcache
	@go test ./... -v