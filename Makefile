.PHONY: run_service
run_service:
	@go run cmd/banner-server/main.go --config=env/local/.env

.PHONY: compose
compose:
	@docker-compose up --build banner-server

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
	@go test ./tests/
