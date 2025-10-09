GOOSE_DRIVER?=postgres
GOOSE_DBSTRING?=postgres://postgres:postgres@localhost:5432/db05?sslmode=disable
GOOSE_MIGRATION_DIR?=migrations

default: help

.PHONY: help
help: pad = 24 # padding for two columns
help:	## Show this help
	@echo "Commands:"
	@fgrep -h "##" $(MAKEFILE_LIST) \
		| fgrep -v fgrep \
		| sed -e 's/^/  /' -e 's/:/ /' -e 's/	//g' \
		| sort -k 1 \
		| grep -v '^  #' \
		| awk -F "#" '{printf ("%s% *s%s\n", $$1, $(pad)-length($$1), "", $$3)}'
	@echo

.PHONY: migrate-status
migrate-status:	## List applied and pending migrations
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" status


.PHONY: migrate-up
migrate-up:		## Create database (if necessary) and migrate to the latest version
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" up

.PHONY: migrate-down
migrate-down:	## Rollback the most recent migration
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" down-to 0

.PHONY: migrate-create
migrate-create:  ## Create new migration files (specify NAME=)
	@if [ -z "$(NAME)" ]; then \
		echo "Error: Migration name is required"; \
		exit 1; \
	fi
	goose -dir $(GOOSE_MIGRATION_DIR) create "$(NAME)" sql
