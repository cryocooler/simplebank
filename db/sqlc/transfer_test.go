package db

import (
	"context"
	"testing"
	"time"

	"github.com/cryocooler/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, a,b Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: a.ID,
		ToAccountID: b.ID,
		Amount: util.RandomMoney(),

	}

	transfer, err := TestQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer

}

func TestCreateTransfer(t *testing.T) {
	a := createRandomAccount(t)
	b := createRandomAccount(t)

	createRandomTransfer(t, a, b)
}

func TestGetTransfer(t *testing.T) {
	a := createRandomAccount(t)
	b := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, a,b)

	transfer2, err := TestQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}


func TestListTransfers(t *testing.T) {
	a := createRandomAccount(t)
	b := createRandomAccount(t)

	for i := 0; i<10; i++ {
		createRandomTransfer(t, a, b)
	}

	arg := ListTransfersParams{
		FromAccountID: a.ID,
		ToAccountID: b.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := TestQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}