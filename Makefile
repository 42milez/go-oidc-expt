PROJECT_NAME := go-oidc-server
VERSION := $(shell git tag --sort=-v:refname | head -n 1)
GITHUB_ID := 42milez

export CI_APP_BUILD_TARGET := dev

.PHONY: $(shell cat $(MAKEFILE_LIST) | awk -F':' '/^[a-z0-9_-]+:/ {print $$1}')

help: Makefile
	@sed -n "s/^##//p" $< | column -t -s ":" |  sed -e "s/^/ /"

# ==================================================
#  Build
# ==================================================

## build: Build a docker image to deploy
build:
	@docker build \
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

## gen: Run generator
gen:
	@go generate ./...

## lint: Run linters
lint:
	@go run -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint run -v --fix
	@vacuum lint -r .vacuum.yml -d app/api/spec/spec.yml

## migrate-apply: Apply migrations
migrate-apply:
	@./scripts/atlas/migrate-apply.sh ${DB_NAMES}

## migrate-diff: Generate migrations
migrate-diff:
ifndef MIGRATION_NAME
	$(error MIGRATION_NAME is required; e.g. make MIGRATION_NAME=xxx migrate-diff)
endif
	@./scripts/atlas/migrate-diff.sh ${MIGRATION_NAME}

## migrate-lint: Run analysis on the migration directory
migrate-lint:
	@./scripts/atlas/migrate-lint.sh ${N_LATEST}

## resolve: Resolve dependencies
resolve:
	@go mod tidy

## seed: Seeding database
seed:
	@go run ./scripts/seed/*.go

## test: Run all tests
test:
	@go test -covermode=atomic -coverprofile=coverage.out `go list ./... | grep -v "/ent" | grep -v "/docs"`

# ==================================================
#  Lima
# ==================================================

lc-create:
	@limactl create --tty=false --name=$(PROJECT_NAME) lima.yml

lc-start:
	@limactl start $(PROJECT_NAME)

lc-stop:
	@limactl stop $(PROJECT_NAME)

lc-delete:
	@limactl delete $(PROJECT_NAME)

# ==================================================
#  Docker
# ==================================================

## up: Create and start containers
up:
	@./scripts/docker/up.sh

## down: Stop and remove containers
down:
	@docker-compose down

## start: Start containers
start:
	@docker-compose start

## stop: Stop containers
stop:
	@docker-compose stop

## destroy: Delete all resources
destroy:
	@docker-compose down --volumes
