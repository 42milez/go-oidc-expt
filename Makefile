PROJECT_NAME := go-oidc-server
VERSION := $(shell git tag --sort=-v:refname | head -n 1)
GITHUB_ID := 42milez

.PHONY: help

help: Makefile
	@sed -n "s/^##//p" $< | column -t -s ":" |  sed -e "s/^/ /"

# ==================================================
#  Build
# ==================================================

.PHONY: build
.PHONY: build-local

## build: Build a docker image to deploy
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

# ==================================================
#  Utility
# ==================================================

.PHONY: benchmark
.PHONY: cleanup-db
.PHONY: cleanup-go
.PHONY: debug
.PHONY: fmt
.PHONY: gen
.PHONY: lint
.PHONY: migrate-apply
.PHONY: migrate-diff
.PHONY: migrate-lint
.PHONY: resolve
.PHONY: seed
.PHONY: test

## benchmark: Run all benchmarks
benchmark:
	@go test -bench . -skip Test.+ -benchmem `go list ./... | grep -v "/ent" | grep -v "/docs"`

## cleanup-db: Clean up database
cleanup-db: export DB_HOST := 127.0.0.1
cleanup-db: export DB_USER := root
cleanup-db: export DB_PORT := 13306
cleanup-db: export DB_NAME := idp
cleanup-db:
	@mysql -h $$DB_HOST -u $$DB_USER -P $$DB_PORT -Nse "show tables" $$DB_NAME | \
		while read table; do \
			[[ $$table == "atlas_schema_revisions" ]] && continue; \
			mysql -h $$DB_HOST -u $$DB_USER -P $$DB_PORT --init-command="SET SESSION FOREIGN_KEY_CHECKS=0;" -e "truncate table $$table" $$DB_NAME; \
		done

## cleanup-go: Clean up caches
cleanup-go:
	@go clean -cache -fuzzcache -testcache

## fmt: Run formatter
fmt:
	@go fmt ./...
	@swag fmt -d app

## gen: Run generator
gen:
	@go generate ./...
	@swag init -d app -o app/docs

## lint: Run linters
lint:
	@golangci-lint run --fix

## migrate-apply: Apply migrations
migrate-apply:
	@./script/atlas/migrate-apply.sh

## migrate-diff: Generate migrations
migrate-diff:
ifndef MIGRATION_NAME
	$(error MIGRATION_NAME is required; e.g. make MIGRATION_NAME=xxx migrate-diff)
endif
	@./script/atlas/migrate-diff.sh ${MIGRATION_NAME}

## migrate-lint: Run analysis on the migration directory
migrate-lint:
	@./script/atlas/migrate-lint.sh ${N_LATEST}

## resolve: Resolve dependencies
resolve:
	@go mod tidy

## seed: Seeding database
seed:
	@go run ./script/seed/main.go

## test: Run all tests
test: export CI := true
test: export DB_PORT := 13306
test: export REDIS_PORT := 16379
test:
	@go test -covermode=atomic -coverprofile=coverage.out `go list ./... | grep -v "/ent" | grep -v "/docs"`

# ==================================================
#  Docker
# ==================================================

.PHONY: down
.PHONY: start
.PHONY: stop
.PHONY: up

## up: Create and start containers
up:
	@docker-compose up -d

## down: Stop and remove containers
down:
	@docker-compose down

## destroy: Destroy all resources
destroy:
	@docker-compose down --volumes

## start: Start containers
start:
	@docker-compose start

## stop: Stop containers
stop:
	@docker-compose stop
