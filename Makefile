PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=cagnotte \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=cagnotted \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=cagnottecli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

include Makefile.ledger
all: build

build: go.sum
		@echo "--> Builing svc and cli"
		@go build -o bin/_svc $(BUILD_FLAGS) ./cmd/cagnotted
		@go build -o bin/_cli $(BUILD_FLAGS) ./cmd/cagnottecli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)