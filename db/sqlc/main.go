package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:plane@localhost:5432/plane?sslmode=disable"
)

var Db *Queries

func Initialize() error {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot Connect to the Database :(", err)
	}

	Db = New(conn)
	return nil
}
