postgresDocker:
	docker run --name go_finance -e POSTGRESQL_USERNAME=gofinance -e POSTGRESQL_PASSWORD=gofinance -e POSTGRESQL_DATABASE=gofinance -p 5432:5432 bitnami/postgresql:latest

migrateUp:
	migrate -path db/migration -database "postgresql://gofinance:gofinance@localhost:5432/gofinance?sslmode=disable" -verbose up 

migrateDown:
	migrate -path db/migration -database "postgresql://gofinance:gofinance@localhost:5432/gofinance?sslmode=disable" -verbose down

server: 
	go run main.go

.PHONY:
	postgresDocker migrateUp migrateDown server