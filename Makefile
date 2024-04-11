.PHONY: run_service
run_service:
	@go run cmd/banner-server/main.go --config=env/.env

.PHONY: compose
compose:
	@docker-compose up --build banner-server