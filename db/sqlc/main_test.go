package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable"
)

var testStore Store

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testStore = NewStore(conn)
	os.Exit(m.Run())
}
