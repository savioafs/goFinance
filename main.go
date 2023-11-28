package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/savioafs/goFinance/api"
	db "github.com/savioafs/goFinance/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://gofinance: gofinance@localhost:5432/gofinance?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api: ", err)
	}
}
