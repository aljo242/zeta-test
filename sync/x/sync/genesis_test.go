package sync_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "sync/testutil/keeper"
	"sync/testutil/nullify"
	"sync/x/sync"
	"sync/x/sync/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		HeaderList: []types.Header{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		HeaderCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SyncKeeper(t)
	sync.InitGenesis(ctx, *k, genesisState)
	got := sync.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.HeaderList, got.HeaderList)
	require.Equal(t, genesisState.HeaderCount, got.HeaderCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
