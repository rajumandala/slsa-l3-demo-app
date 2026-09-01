.PHONY: build test run clean

VERSION ?= dev

build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/hello ./cmd/hello

test:
	go test ./...

run: build
	./bin/hello

clean:
	rm -rf bin
