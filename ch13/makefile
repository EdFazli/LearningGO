.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint:
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
	shadow ./...
.PHONY:vet

build: vet
	go build tests.go
.PHONY:build