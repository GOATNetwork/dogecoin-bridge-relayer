# Makefile for nudex_voter

VERSION := $(shell git describe --tags --always)
GITREV := $(shell git rev-parse --short HEAD)
GITBRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DATE := $(shell LANG=US date +"%a, %d %b %Y %X %z")

GOMOD := github.com/nuvosphere/nudex-voter

LDFLAGS += -X '$(GOMOD)/version.Version=$(VERSION)'
LDFLAGS += -X '$(GOMOD)/version.GitRev=$(GITREV)'
LDFLAGS += -X '$(GOMOD)/version.GitBranch=$(GITBRANCH)'
LDFLAGS += -X '$(GOMOD)/version.BuildDate=$(DATE)'

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=nudex-voter

# Build binary
all: build

build:
	$(GOBUILD) -ldflags "all=$(LDFLAGS)" -o $(BINARY_NAME) -v ./cmd
	rm -rf ./test/node1/nudex-voter
	rm -rf ./test/node2/nudex-voter
	rm -rf ./test/node3/nudex-voter
	cp nudex-voter ./test/node1/
	cp nudex-voter ./test/node2/
	cp nudex-voter ./test/node3/

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf ./test/node1/data/db/*.db
	rm -rf ./test/node1/data/db/*.db-*
	rm -rf ./test/node2/data/db/*.db
	rm -rf ./test/node2/data/db/*.db-*
	rm -rf ./test/node3/data/db/*.db
	rm -rf ./test/node3/data/db/*.db-*
	rm -rf ./test/node1/data/db/chain_id_*
	rm -rf ./test/node2/data/db/chain_id_*
	rm -rf ./test/node3/data/db/chain_id_*
	rm -rf ./test/node1/nudex-voter
	rm -rf ./test/node2/nudex-voter
	rm -rf ./test/node3/nudex-voter
.PHONY: test
test: ## go test
	go test  -short -v ./...

deps:
	$(GOGET) -u ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd && ./$(BINARY_NAME)

docker-build-all:
	docker buildx build --platform linux/amd64,linux/arm64 -t nuvosphere/nudex-voter:latest --push .

docker-build:
	docker buildx build --platform linux/amd64 -t nuvosphere/nudex-voter:latest --load .

docker-build-x:
	docker buildx build --platform linux/arm64 -t nuvosphere/nudex-voter:latest --load .

.PHONY: fix
fix: ## auto fix code
	golangci-lint run --fix ./...

.PHONY: dep
tool: ## Install dep tool
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install github.com/daixiang0/gci@latest
	go install github.com/bombsimon/wsl/v4/cmd/wsl@latest
	go install github.com/tetafro/godot/cmd/godot@latest
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/4meepo/tagalign/cmd/tagalign@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

.PHONY: abi
abi: ## generate abi go file
	abigen --abi internal/layer2/abis/erc20.json --pkg contracts --type ERC20 --out internal/layer2/contracts/erc20.go
	abigen --abi internal/layer2/abis/ierc1271.json --pkg contracts --type IERC1271 --out internal/layer2/contracts/ierc1271.go
	abigen --abi internal/layer2/abis/erc721.json --pkg contracts --type ERC721 --out internal/layer2/contracts/erc721.go
	abigen --abi internal/layer2/abis/erc1155.json --pkg contracts --type ERC1155 --out internal/layer2/contracts/erc1155.go
	abigen --abi internal/layer2/abis/multicall3.json --pkg contracts --type Multicall3 --out internal/layer2/contracts/multicall3.go
	abigen --abi internal/layer2/abis/AccountManager.json --pkg contracts --type AccountManagerContract --out internal/layer2/contracts/account_manager.go
	abigen --abi internal/layer2/abis/TaskManager.json --pkg contracts --type TaskManagerContract --out internal/layer2/contracts/task_manager.go
	abigen --abi internal/layer2/abis/AssetHandler.json --pkg contracts --type AssetHandlerContract --out internal/layer2/contracts/asset_handler.go
	abigen --abi internal/layer2/abis/ParticipantManager.json --pkg contracts --type ParticipantManagerContract --out internal/layer2/contracts/participant_manager.go
	abigen --abi internal/layer2/abis/VotingManager.json --pkg contracts --type VotingManagerContract --out internal/layer2/contracts/voting_manager.go
	abigen --abi internal/layer2/abis/DepositManager.json --pkg contracts --type DepositManagerContract --out internal/layer2/contracts/deposit_manager.go
	abigen --abi internal/layer2/abis/TaskPayload.json --pkg contracts --type TaskPayloadContract --out internal/layer2/contracts/task_payload.go
	abigen --abi internal/layer2/abis/codec/Codec.json --pkg codec --type VoterCodec --out internal/layer2/contracts/codec/codec.go
	abigen --abi internal/layer2/abis/Inscription.json --pkg contracts --type InscriptionContract --out internal/layer2/contracts/inscription.go

contract:
	solc --pretty-json --optimize --abi  --overwrite  -o internal/layer2/abis/codec/    ./internal/layer2/contracts/Codec.sol
	mv internal/layer2/abis/codec/Codec.abi internal/layer2/abis/codec/Codec.json
	solc --pretty-json --optimize --abi  --overwrite  -o internal/layer2//contracts/    ./internal/layer2/contracts/TaskPayload.sol
	mv internal/layer2/contracts/TaskPayload.abi internal/layer2/abis/TaskPayload.json
	

.PHONY: ci
ci: abi fix build test 
	echo $? && echo "success!"

.PHONY: lint
lint: ## Runs the linter
	golangci-lint run ./...

.PHONY: help
help: ## Prints this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	
.PHONY: all build clean test deps run docker-build