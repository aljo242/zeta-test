build:
	@go build .

run:
	@go run .

run-chain:
	@cd sync && ignite chain serve --verbose

run-all: run run-chain

lint:
	@echo "--> Running linter"
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --out-format=tab

lint-fix:
	@echo "--> Running linter"
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --fix --out-format=tab --issues-exit-code=0

test:
	@go test ./...
	@cd sync && go test ./..