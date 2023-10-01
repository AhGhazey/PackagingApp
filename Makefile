PROJECT=github/ahmedghazey/packaging
export CURRENT_DIR=$(shell pwd)

all: mod fmt build

mod:
	go mod tidy
	go mod vendor

fmt:
	@echo "Running go fmt"
	go fmt $(PROJECT)/...

build:
	@echo "Building"
	go build -o packaging cmd/api/main.go

run:
	@echo "Running:"
	go run cmd/api/main.go

test:
	@echo "Running tests"
	go clean -testcache
	go test -coverprofile=coverage.out --race -p 1 ./...

coverage:
	go tool cover -html=./coverage.out


swag:
	@echo "Generating swagger files .."
	swag init --dir ./cmd/api/,./internal/http/rest/ --markdownFiles ./README.md  --output ./docs --parseDependency

.PHONY: all mod fmt build run test coverage swag