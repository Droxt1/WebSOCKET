
POSTGRES_CONTAINER_NAME = WebSocket-Postgres
POSTGRES_PORT = 5433
POSTGRES_USER = root
POSTGRES_PASSWORD = password
POSTGRES_DB = go-chat
MIGRATIONS_DIR = db/migrations

dockerup:
	docker-compose up

dockerdown:
	docker-compose down

startcontainer:
	docker start $(POSTGRES_CONTAINER_NAME)

stopcontainer:
	docker stop $(POSTGRES_CONTAINER_NAME)

postgres:
	docker exec -it $(POSTGRES_CONTAINER_NAME) psql

psqlpostgres:
	psql -h localhost -p $(POSTGRES_PORT) -U $(POSTGRES_USER) -d $(POSTGRES_DB)

createdb:
	docker exec -it $(POSTGRES_CONTAINER_NAME) createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_DB)

dropdb:
	docker exec -it $(POSTGRES_CONTAINER_NAME) dropdb $(POSTGRES_DB)

createmigration:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) add_users_table

migrateup:
	migrate -path $(MIGRATIONS_DIR) -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" -verbose up

migratedown:
	migrate -path $(MIGRATIONS_DIR) -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" -verbose down

vet:
	go vet ./...

test:
	go test -v ./...

golint:
	golangci-lint run

fmt:
	go fmt ./...

build:
	go build -o app ./cmd

run:
	./app

# Declare targets as phony (no actual files with these names)
.PHONY: dockerup dockerdown startcontainer stopcontainer postgres psqlpostgres createdb dropdb createmigration migrateup migratedown vet test golint fmt build run
