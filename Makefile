.PHONY: run_service
run_service:
	@go run cmd/banner-server/main.go --config=env/local/.env

.PHONY: staticcheck
staticcheck:
	@go vet -vettool=$(which staticcheck -f) ./...

.PHONY: tests
tests:
	@go test ./tests/end_to_end/post_banner_test.go
	@go test ./tests/end_to_end/get_banners_test.go
	@go test ./tests/end_to_end/get_user_banner_test.go
	@go test ./tests/end_to_end/patch_banners_test.go
	@go test ./tests/end_to_end/delete_banner_test.go

.PHONY: stress_test
stress_test:
	@go test ./tests/1000RPS/stress.go
	@go test ./tests/2000RPS/stress.go

.PHONY: gen
gen:
	@go run ./tests/gen/generate_data.go

.PHONY: compose-build-run
compose-build-run:
	@docker-compose up --build banner-server

.PHONY: compose-build
compose-build:
	@docker-compose build banner-server

.PHONY: compose-run
compose-run:
	@docker-compose up banner-server

.PHONY: compose-down
compose-down:
	@docker-compose down