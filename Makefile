MIGRATIONS_PATH = ./db/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

# Define variables
POSTGRESQL_URL = postgres://library:librarypassword@localhost:5432/library_db?sslmode=disable

# Define the default goal
.DEFAULT_GOAL := help

# Help target
help:
	@echo "Available targets:"
	@echo "  migrate-up   : Run database migrations"
	@echo "  migrate-down : Rollback database migrations"

# Migration targets
migrate-up:
	@echo "Running database migrations..."
	@migrate -database "${POSTGRESQL_URL}" -path db/migrations up

migrate-down:
	@echo "Rolling back database migrations..."
	@migrate -database "${POSTGRESQL_URL}" -path db/migrations down

# Ensure that the POSTGRESQL_URL is exported for use in commands
export POSTGRESQL_URL

.PHONY: help migrate-up migrate-down