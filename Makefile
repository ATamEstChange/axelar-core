PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=axelar \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=axelard \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=axelarcli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

.PHONY: all
all: generate build docker-image docker-image-debug

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

# Uncomment when you have some tests
# test:
# 	@go test -mod=readonly $(PACKAGES)
.PHONY: lint
# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

# Build the project with release flags
.PHONY: build
build: go.sum
		go build -o ./bin/axelard -mod=readonly $(BUILD_FLAGS) ./cmd/axelard
		go build -o ./bin/axelarcli -mod=readonly $(BUILD_FLAGS) ./cmd/axelarcli
		go build -o ./bin/vald -mod=readonly $(BUILD_FLAGS) ./cmd/vald


# Build the project with debug flags
.PHONY: debug
debug: go.sum
		go build -o ./bin/axelard -mod=readonly $(BUILD_FLAGS) -gcflags="all=-N -l" ./cmd/axelard
		go build -o ./bin/axelarcli -mod=readonly $(BUILD_FLAGS) -gcflags="all=-N -l" ./cmd/axelarcli
		go build -o ./bin/vald -mod=readonly $(BUILD_FLAGS) -gcflags="all=-N -l" ./cmd/vald

# Build axelarcli with release flags for alpine architecture
.PHONY: alpine-axelarcli
alpine-axelarcli: go.sum
		GOOS=linux GOARCH=amd64 go build -o ./bin/axelarcli -mod=readonly $(BUILD_FLAGS) ./cmd/axelarcli

# Build a release image
.PHONY: docker-image
docker-image:
	@DOCKER_BUILDKIT=1 docker build --ssh default -t axelar/core .

# Build a docker image that is able to run dlv and a debugger can be hooked up to
.PHONY: docker-image-debug
docker-image-debug:
	@DOCKER_BUILDKIT=1 docker build --ssh default -t axelar/core-debug -f ./Dockerfile.debug .

# Install all generate prerequisites
.Phony: prereqs
prereqs:
ifeq (which moq,)
	go get github.com/matryer/moq
endif
ifeq (which mdformat,)
	pip3 install mdformat
endif
ifeq (which protoc,)
	@echo "Please install protoc for grpc (https://grpc.io/docs/languages/go/quickstart/)"
endif

# Run all the code generators in the project
.PHONY: generate
generate:
	go generate -x ./...


.PHONE: tofnd-client
tofnd-client:
	@echo -n Generating protobufs...
	@protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative x/tss/tofnd/tofnd.proto
	@echo done

