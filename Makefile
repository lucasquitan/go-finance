createDB:
	docker exec -it postgres createdb --username=postgres go_finance

postgres:
	docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:14-alpine

migrationUP:
	migrate -path src/database/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

migrationDrop:
	migrate -path src/database/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose drop


.PHONY: createDB postgres
