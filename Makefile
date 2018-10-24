BIN_NAME = docker-discovery

all: build run

build: 
	@go build -o bin/${BIN_NAME} ./cmd/${BIN_NAME}

get-deps:
	@dep ensure -v

run:
	@./bin/docker-discovery
