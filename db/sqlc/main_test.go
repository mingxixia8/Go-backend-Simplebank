package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" //_ means keep this import
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connnect to database:", err)
	}
	// if connec to db, run unit test
	testQueries = New(conn)
	os.Exit(m.Run())
}
