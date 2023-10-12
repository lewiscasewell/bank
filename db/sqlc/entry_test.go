package db

import (
	"context"
	"testing"

	"github.com/lewiscasewell/bank/util"
)

func createRandomEntry(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testStore.CreateEntry(context.Background(), arg)
	if err != nil {
		t.Error(err)
	}

	if entry.AccountID != arg.AccountID {
		t.Errorf("account id is different")
	}
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}
