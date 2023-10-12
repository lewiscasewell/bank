package db

import (
	"context"
	"testing"
	"time"

	"github.com/lewiscasewell/bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account, amount int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
	}

	transfer, err := testStore.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func getAccountById(t *testing.T, id int64) Account {
	acc, err := testStore.GetAccount(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	return acc
}

func TestCreateTransfer(t *testing.T) {
	amount := util.RandomMoney()
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2, amount)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := util.RandomMoney()

	randomTransfer := createRandomTransfer(t, account1, account2, amount)
	transfer, err := testStore.GetTransfer(context.Background(), randomTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, randomTransfer.ID, transfer.ID)
	require.Equal(t, randomTransfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, randomTransfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, randomTransfer.Amount, transfer.Amount)
	require.WithinDuration(t, randomTransfer.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	account3 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2, util.RandomMoney())
		createRandomTransfer(t, account1, account3, util.RandomMoney())
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testStore.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account2.ID)
	}
}
