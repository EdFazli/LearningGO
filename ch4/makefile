.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint:
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build controlStructures.go
.PHONY:build