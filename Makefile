PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := v0.2.2-dev
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=uro-token \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=uro-token \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

install: go.sum
	@echo "--> Installing Uro token"
	@rm -rf build
	@go build  $(BUILD_FLAGS) ./cmd/uro
	@mkdir build
	@mv uro build/uro
