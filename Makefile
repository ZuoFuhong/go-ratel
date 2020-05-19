BINARY = ratel
VERSION = 1.0.0
GOOS = darwin
GOARCH = amd64

build:
	@export GOOS=${GOOS}; \
  	export GOARCH=${GOARCH}; \
  	go build -o ${BINARY}

test: build
	@./$(BINARY) -h 39.105.65.8 -p 1024

.PHONY: build test
