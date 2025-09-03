# Load .env variables
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

DB_URL = postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_PATH = migrations

# -----------------------
# Migration commands
# -----------------------

# Run all up migrations
migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

# Rollback last migration
migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

# Force version (fix dirty db)
migrate-force:
	@echo "Enter version to force:"; \
	read version; \
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $$version

# Reset database (down all + up)
migrate-reset:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

# -----------------------
# Create new migration file
# -----------------------
create-migration:
	@echo "Enter migration name:"; \
	read name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name

.PHONY: migrate-up migrate-down migrate-force migrate-reset create-migration