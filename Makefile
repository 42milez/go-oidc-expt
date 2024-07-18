PROJECT_NAME := go-oidc-expt
VERSION := $(shell git tag --sort=-v:refname | head -n 1)
GITHUB_ID := 42milez

.PHONY: $(shell cat $(MAKEFILE_LIST) | awk -F':' '/^[a-z0-9_-]+:/ {print $$1}')

help: Makefile
	@sed -n "s/^##//p" $< | column -t -s ":" |  sed -e "s/^/ /"

define run
	cmd=$1; \
	echo "▶ RUN: $${cmd}"; eval "$${cmd}"
endef

# ==================================================
#  Build
# ==================================================

## build: Build a docker image to deploy
build:
	@cmd="docker build --no-cache -f docker/app/Dockerfile -t ${GITHUB_ID}/${PROJECT_NAME}:${VERSION} --build-arg VERSION=${VERSION} --target deploy ."; \
		$(call run,"$${cmd}")

## build-local: Build docker images
build-local:
	@cmd="docker compose build --no-cache"; \
		$(call run,"$${cmd}")

# ==================================================
#  Utility
# ==================================================

## benchmark: Run all benchmarks
benchmark:
	@cmd='go test -bench . -skip Test.+ -benchmem `go list ./... | grep -v "/ent" | grep -v "/docs"`'; \
		$(call run,"$${cmd}")

## cleanup-db: Clean up database
cleanup-db: export DB_HOST := 127.0.0.1
cleanup-db: export DB_USER := root
cleanup-db: export DB_PORT := 13306
cleanup-db: export DB_NAME := idp
cleanup-db:
	@mysql -h $$DB_HOST -u $$DB_USER -P $$DB_PORT -Nse "show tables" $$DB_NAME | \
		while read table; do \
			[[ $$table == "atlas_schema_revisions" ]] && continue; \
			cmd="mysql -h $$DB_HOST -u $$DB_USER -P $$DB_PORT --init-command='SET SESSION FOREIGN_KEY_CHECKS=0;' -e 'truncate table $$table' $$DB_NAME"; \
			$(call run,"$${cmd}"); \
		done

## cleanup-go: Clean up caches
cleanup-go:
	@cmd="go clean -cache -fuzzcache -testcache"; \
		$(call run,$${cmd})

## fmt: Run formatter
fmt:
	@commands=( \
		"go run -mod=mod golang.org/x/tools/cmd/goimports -w ./cmd" \
		"go run -mod=mod golang.org/x/tools/cmd/goimports -w ./pkg" \
		"go run -mod=mod golang.org/x/tools/cmd/goimports -w ./tools" \
		"go run -mod=mod golang.org/x/tools/cmd/goimports -w ./scripts" \
		"go fmt ./..." \
	); \
		for cmd in "$${commands[@]}"; do \
			$(call run,$${cmd}); \
		done

## gen: Run generator
gen:
	@cmd="go generate ./..."; \
		$(call run,$${cmd})

## lint: Run linters
lint:
	@commands=( \
		"go run -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint run --fix" \
		"vacuum lint -r .vacuum.yml -d cmd/api/spec/spec.yml" \
	); \
		for cmd in "$${commands[@]}"; do \
			$(call run,$${cmd}); \
		done

## migrate-apply: Apply migrations
migrate-apply:
ifndef SERVICE
	$(error SERVICE is required; e.g. make migrate-apply SERVICE=*** DATABASE=***)
else ifndef DATABASE
	$(error DATABASE is required; e.g. make migrate-apply SERVICE=*** DATABASE=***)
endif
	@cmd="./scripts/atlas/migrate.sh apply --service $${SERVICE} --database $${DATABASE}"; \
		$(call run,$${cmd})

## migrate-diff: Generate migrations
migrate-diff:
ifndef MIGRATION_NAME
	$(error MIGRATION_NAME is required; e.g. make migrate-diff MIGRATION_NAME=***)
endif
	@cmd="./scripts/atlas/migrate.sh diff --migration-name $${MIGRATION_NAME}"; \
		$(call run,$${cmd})

## migrate-lint: Run analysis on the migration directory
migrate-lint:
ifdef LATEST
	@cmd="./scripts/atlas/migrate.sh lint --latest $${LATEST}"; \
		$(call run,$${cmd})
else
	@cmd="./scripts/atlas/migrate.sh lint"; \
		$(call run,$${cmd})
endif

## resolve: Resolve dependencies
resolve:
	@cmd="go mod tidy"; \
		$(call run,$${cmd})

## seed: Seeding database
seed:
	@cmd="go run ./scripts/seed/*.go"; \
		$(call run,$${cmd})

## test: Run all tests
test:
	@go clean -testcache
	@cmd='go test -covermode=atomic -coverprofile=coverage.out `go list ./... | grep -v "/ent" | grep -v "/docs" | grep -v "/scripts"`'; \
		$(call run,$${cmd})

# ==================================================
#  Lima
# ==================================================

# lc-create: Create virtual machine with Lima
#lc-create:
#	@cmd="limactl create --tty=false --name=$(PROJECT_NAME) lima.yml"; \
#		$(call run,$${cmd})

# lc-start: Start virtual machine
#lc-start:
#	@cmd="limactl start $(PROJECT_NAME)"; \
#		$(call run,$${cmd})

# lc-stop: Stop virtual machine
#lc-stop:
#	@cmd="limactl stop $(PROJECT_NAME)"; \
#		$(call run,$${cmd})

# lc-delete: Delete virtual machine
#lc-delete:
#	@cmd="limactl delete $(PROJECT_NAME)"; \
#		$(call run,$${cmd})

# lc-restart: Restart virutal machine
#lc-restart:
#	@cmd="limactl stop $(PROJECT_NAME) && limactl start $(PROJECT_NAME)"; \
#		$(call run,$${cmd})

# ==================================================
#  Docker
# ==================================================

## up: Create and start containers
up:
	@cmd="./scripts/docker/up.sh"; \
		$(call run,$${cmd})

## down: Stop and remove containers
down:
	@cmd="docker compose down"; \
		$(call run,$${cmd})

## start: Start containers
start:
	@cmd="docker compose start"; \
		$(call run,$${cmd})

## stop: Stop containers
stop:
	@cmd="docker compose stop"; \
		$(call run,$${cmd})

## restart: Restart containers
restart:
	@cmd="make down && make up"; \
		$(call run,$${cmd})

## destroy: Delete all resources
destroy:
	@cmd="docker compose down --volumes"; \
		$(call run,$${cmd})
