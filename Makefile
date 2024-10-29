# Create the database directly from the simplebank_app container
createdb:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d postgres -c "CREATE DATABASE $(DB_NAME);"

# Run migrations directly from the simplebank_app container
migrateup:
	migrate -path=/app/db/migration -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up

# Rollback the most recent migration
migratedown:
	migrate -path=/app/db/migration -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable down 1

# Drop the database directly from the simplebank_app container
dropdb:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname='$(DB_NAME)';"
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d postgres -c "DROP DATABASE $(DB_NAME);"

# Generate SQL code using sqlc
sqlc:
	sqlc generate

# Run tests
test:
	go test -v -cover ./...

# Generate mocks
mock:
	mockgen -package mockdb -destination db/mock/store.go simple_bank/db/sqlc Store

# Run the server in development mode
.PHONY: createdb dropdb migrateup migratedown sqlc test mock
