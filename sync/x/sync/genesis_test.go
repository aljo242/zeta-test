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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SyncKeeper(t)
	sync.InitGenesis(ctx, *k, genesisState)
	got := sync.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
