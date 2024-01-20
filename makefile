.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint:
	staticcheck ./...
.PHONY:lint

imports:
	gopls imports -w main.go
.PHONY:imports

build: fmt
	go build -ldflags='-s -w' -o ./bin/dataru main.go
	./bin/dataru -config ./default_config.yaml
.PHONY:build

checkrace:
	go run --race main.go
.PHONY:checkrace

debug:
	go build -gcflags='all=-N -l' -o bin/debug main.go
.PHONY:debug

test:
	go test -v ./...
.PHONY:test

bench:
	go test -bench=. ./... -benchmem
