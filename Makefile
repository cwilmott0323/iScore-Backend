.PHONY: build test help fmt

default: help

build: fmt test ## Builds an executable
	GO111MODULE=on go build

test: ## Runs go test with coverage
	GO111MODULE=on go test ./... -cover

fmt: ## Verifies all files have been `gofmt`ed
	@echo "+ $@"
	@gofmt -s -l . | grep -v '.pb.go:' | grep -v vendor | tee /dev/stderr

install: test ## Installs the executable or package
	GO111MODULE=on go install

compile: fmt test ## Compile for all OS
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o iScore-api-freebsd-386
	env GOOS=linux GOARCH=amd64 go build -o iScore-api-linux-amd64
	GOOS=windows GOARCH=386 go build -o iScore-api-windows-386

zip: ## zip binary and .env file for lambda deployment
	zip -j iScore.zip iScore-api-linux-amd64 .env


help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'