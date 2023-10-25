postgresinit:
	docker run --name websocketpostgres1 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

startpostgres:
	sudo docker start websocketpostgres1

postgres:
	docker exec -it websocketpostgres1 psql

psqlpostgres:
	psql -h localhost -p 5433 -U root -d go-chat

createdb:
	docker exec -it websocketpostgres1 createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it websocketpostgres1 dropdb go-chat

createmigration:
	migrate create -ext sql -dir db/migrations add_users_table

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down

vet:
	go vet ./...

test:
	go test -v ./...

golint:
	golangci-lint run

fmt:
	go fmt ./...



.PHONY: postgresini postgres createdb dropdb migrateup migratedown, vet, test, golint, fmt, startpostgres, psqlpostgres, createmigration