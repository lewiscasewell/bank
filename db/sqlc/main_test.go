package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lewiscasewell/bank/util"
)

const (
	dbSource = "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testStore = NewStore(conn)
	os.Exit(m.Run())
}
