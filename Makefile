.PHONY: build clean tool lint help

all: build

docker_build:
	docker build -t my-golang-app .

docker_run:
	docker run -p 8080:8080 -it --rm --name my-running-app my-golang-app

build:
	@go build -v .

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	del go_test.exe
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"