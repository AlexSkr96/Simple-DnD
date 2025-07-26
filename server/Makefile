.DEFAULT_GOAL := help

GO       ?= go
GOFLAGS  ?=
LOCAL_BIN = $(CURDIR)/bin
PROJECT_NAME = "github.com/AlexSkr96/Simple-DnD"

# Include environment variables from local .env file if it exists
-include .env

DB_USER ?= user
DB_PWD ?= pwd
DB_NAME ?= dnd
DB_PORT ?= 5432
DB_HOST ?= localhost
DB_CONN ?=host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PWD) dbname=$(DB_NAME) sslmode=disable
GOOSE_DIR = "-dir=sql/schema"
GOOSE_TABLE = "-table=dnd_goose_db_version"

include bin-deps.mk

.PHONY: lint
lint: $(GOLINT_BIN) ## lints the code
	@for f in $$(git status --porcelain | awk 'match($$1, "M|A"){print $$2}' | grep '.go$$'); do \
		DIR="$$(dirname "$${f}")";\
		$(GOLINT_BIN) run --fix --color=always --timeout 4m $$DIR; \
	done

.PHONY: lint-all
lint-all: $(GOLINT_BIN) ## lints all the code
	$(GOLINT_BIN) run --fix --color=always --timeout 4m;

.PHONY: format
format: $(GOFUMPT_BIN) $(IMPREVISER_BIN) ## formats the code and also imports order
	@echo "Formatting..."
	@for f in $$(git status --porcelain | awk 'match($$1, "M|A|\\?"){print $$2}' | grep '.go$$'); do \
		$(IMPREVISER_BIN) -project-name $(PROJECT_NAME) $$f; \
		$(GOFUMPT_BIN) -l -w -extra $$f; \
	done

.PHONY: test
test: ## runs the tests
	go test ./...

.PHONY: db-mi-up
db-mi-up: $(GOOSE_BIN) ## executes all new db migrations. Requires DB_CONN
	@PATH="$(LOCAL_BIN):$(PATH)"; \
	GOOSE_DRIVER=postgres GOOSE_DBSTRING="$(DB_CONN)" $(GOOSE_BIN) $(GOOSE_DIR) $(GOOSE_TABLE) up

.PHONY: db-mi-down
db-mi-down: $(GOOSE_BIN) ## rollbacks the last db migration. Requires DB_CONN
	@PATH="$(LOCAL_BIN):$(PATH)"; \
	GOOSE_DRIVER=postgres GOOSE_DBSTRING="$(DB_CONN)" $(GOOSE_BIN) $(GOOSE_DIR) $(GOOSE_TABLE) down

.PHONY: db-mi-status
db-mi-status: $(GOOSE_BIN) ## shows the status of db migrations. Requires DB_CONN
	@PATH="$(LOCAL_BIN):$(PATH)"; \
 	GOOSE_DRIVER=postgres GOOSE_DBSTRING="$(DB_CONN)" $(GOOSE_BIN) $(GOOSE_DIR) $(GOOSE_TABLE) status

.PHONY: db-mi-create
db-mi-create: $(GOOSE_BIN) ## creates a new db migration. Example: `make db-mi-create NAME=new_migration_name`
	@PATH="$(LOCAL_BIN):$(PATH)"; \
	GOOSE_DRIVER=postgres $(GOOSE_BIN) $(GOOSE_DIR) create $(NAME) sql

compose-build:
	docker-compose -f docker-compose.yml build

compose-up:
	docker-compose -f docker-compose.yml up

compose-down:
	docker-compose -f docker-compose.yml down --remove-orphans

.PHONY: help
help: ## shows this help
	@grep --no-filename -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
