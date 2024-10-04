postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

to run psql console:
	docker exec -it postgres16 psql -U root -d simple_bank

.PHONY: createdb


docker exec -it postgres16 psql -U root simple_bank	TO access the db console without
going through the container shell