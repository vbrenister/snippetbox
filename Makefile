server:
	@go run ./cmd/web


migrateup:
	@migrate -database "postgres://snippets:snippets@localhost:5533/snippets?sslmode=disable" -path db/migrations --verbose up 

migratedown:
	@migrate -database "postgres://snippets:snippets@localhost:5533/snippets?sslmode=disable" -path db/migrations --verbose down

createdb:
	@docker run --name=snippets-db -d --env POSTGRES_DB=snippets --env POSTGRES_PASSWORD=snippets --env POSTGRES_USER=snippets -p 5533:5432 postgres:11.12

dropdb:
	@docker stop snippets-db && docker rm snippets-db 

test:
	@go test -v ./...
	
.PHONY: server migrateup migratedown createdb dropdb test