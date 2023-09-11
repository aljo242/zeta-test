package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "sync/testutil/keeper"
	"sync/testutil/nullify"
	"sync/x/sync/keeper"
	"sync/x/sync/types"
)

func createNHeader(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Header {
	items := make([]types.Header, n)
	for i := range items {
		items[i].Id = keeper.AppendHeader(ctx, items[i])
	}
	return items
}

func TestHeaderGet(t *testing.T) {
	keeper, ctx := keepertest.SyncKeeper(t)
	items := createNHeader(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHeader(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHeaderRemove(t *testing.T) {
	keeper, ctx := keepertest.SyncKeeper(t)
	items := createNHeader(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHeader(ctx, item.Id)
		_, found := keeper.GetHeader(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHeaderGetAll(t *testing.T) {
	keeper, ctx := keepertest.SyncKeeper(t)
	items := createNHeader(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHeader(ctx)),
	)
}

func TestHeaderCount(t *testing.T) {
	keeper, ctx := keepertest.SyncKeeper(t)
	items := createNHeader(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHeaderCount(ctx))
}
