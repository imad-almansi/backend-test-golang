IMAGE_NAME := facts
VERSION := latest

.PHONY: build clean image run test

build:
	go build -o bin/facts ./*.go

clean:
	docker compose down
	rm bin/*

image:
	docker build --tag ${IMAGE_NAME}:${VERSION} .

run:
	docker compose up

test:
	go test -v ./...