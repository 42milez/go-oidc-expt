PROJECT_NAME := go-oidc-server
VERSION := $(shell git tag --sort=-v:refname | head -n 1)
GITHUB_ID := 42milez

.PHONY: help build build-local clean debug down fmt lint migrate resolve start stop test up

help: Makefile
	@sed -n "s/^##//p" $< | column -t -s ":" |  sed -e "s/^/ /"

#  Build
# --------------------------------------------------

## build: Build docker images to deploy
build:
	@docker buildx build \
		--no-cache \
		-f docker/app/Dockerfile \
		-t ${GITHUB_ID}/${PROJECT_NAME}:${VERSION} \
		--build-arg VERSION=${VERSION} \
		--target deploy \
		.

## build-local: Build docker images
build-local:
	@docker-compose build --no-cache

#  Utility
# --------------------------------------------------

## clean: Clean up caches
clean:
	@go clean -cache -fuzzcache -testcache

## debug: Start containers with debugger
debug:
	@docker-compose -f docker-compose.yml -f docker/docker-compose.debug.yml up --force-recreate -d

## fmt: Run formatter
fmt:
	@go fmt ./...

## lint: Run linters
lint:
	@golangci-lint run --fix

## migrate: Run migration
migrate:
	@echo "not implemented"

## resolve: Resolve dependencies
resolve:
	@go mod tidy

## test: Run all tests
test:
	@go test -covermode=atomic -coverprofile=coverage.out ./...

#  Docker
# --------------------------------------------------

## up: Create and start containers
up:
	@docker-compose up --force-recreate -d

## down: Stop and remove containers
down:
	@docker-compose down

## start: Start containers
start:
	@docker-compose start

## stop: Stop containers
stop:
	@docker-compose stop
