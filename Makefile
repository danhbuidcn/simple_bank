# Tạo cơ sở dữ liệu trực tiếp từ container go_app
createdb:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d postgres -c "CREATE DATABASE $(DB_NAME);"

# Chạy các migration trực tiếp từ container go_app
migrateup:
	migrate -path=/app/db/migrations -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up

# Rollback 1 migrate gần nhất
migratedown:
	migrate -path=/app/db/migrations -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable down 1

# Xóa cơ sở dữ liệu trực tiếp từ container go_app
dropdb:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname='$(DB_NAME)';"
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d postgres -c "DROP DATABASE $(DB_NAME);"

sqlc:
	sqlc generate

.PHONY: createdb dropdb migrateup migratedown sqlc
