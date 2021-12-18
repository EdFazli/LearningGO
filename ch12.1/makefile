.DEFAULT_GOAL := build

.PHONY: build run-ok run-cancel

build:
	go build

run-ok: build
	./ch12.1 false

run-cancel: build
	./ch12.1 true