POSTGRES_CONTAINER_NAME=postgres12
DB_NAME=automation_engine
DB_USER=root
DB_PASSWORD=root
DB_PORT=5432
MIGRATION_PATH=database/migration

postgres:
	docker run --name $(POSTGRES_CONTAINER_NAME) -p $(DB_PORT):$(DB_PORT) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:12-alpine

createdb:
	docker exec -it $(POSTGRES_CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(POSTGRES_CONTAINER_NAME) dropdb $(DB_NAME)

migrateup:
	migrate -path $(MIGRATION_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path $(MIGRATION_PATH) -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown test server
