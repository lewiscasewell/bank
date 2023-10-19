DB_CONTAINER=postgres15
DB_USER=postgres
DB_PASSWORD=postgres
DB_PORT=5433
DB_NAME=bank

postgres:
	docker run --name $(DB_CONTAINER) --network bank-network -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p 5433:5432 -d postgres

createdb:
	docker exec -it $(DB_CONTAINER) createdb -U $(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(DB_CONTAINER) dropdb -U $(DB_USER) $(DB_NAME)

migrateup:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/lewiscasewell/bank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1