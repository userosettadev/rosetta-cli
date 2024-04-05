VERSION=$(shell git describe --always --tags | cut -d "v" -f 2)
LINKER_FLAGS=-s -w -X github.com/userosettadev/rosetta-cli/build.Version=${VERSION}
GOLANGCILINT_VERSION=v1.57.2

.PHONY: build
build:  ## Go build
	@echo "==> Building rosetta binary"
	go build -ldflags "$(LINKER_FLAGS)" -o ./bin/rosetta $(MCLI_SOURCE_FILES)

.PHONY: deps
deps:  ## Download go module dependencies
	@echo "==> Installing go.mod dependencies..."
	go mod download
	go mod tidy

.PHONY: test
test:  ## Go recompile and test with coverage atomic
	go test ./... -covermode=atomic

.PHONY: update
update:  ## Update go dependencies
	go get -u

.PHONY: lint
lint:  ## Run linter
	golangci-lint run

.PHONY: devtools
devtools:  ## Install dev tools
	@echo "==> Installing dev tools..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCILINT_VERSION)

.PHONY: dockerbuild
dockerbuild:  ## Docker build
	docker build -t rosetta .

.PHONY: protoc
protoc:  ## Protocol buffer build
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
    internal/api/create_oas.pb
