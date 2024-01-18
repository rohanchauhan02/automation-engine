postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root automation_engine

dropdb:
	docker exec -it postgres12 dropdb automation_engine

migrateup:
	migrate -path db/migration -database "localhost://root:root@localhost:5432/automation_engine?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "localhost://root:root@localhost:5432/automation_engine?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY postgres createdb dropdb migratedown migrateup server test
