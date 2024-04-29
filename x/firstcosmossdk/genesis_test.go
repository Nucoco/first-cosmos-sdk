package firstcosmossdk_test

import (
	"testing"

	keepertest "first-cosmos-sdk/testutil/keeper"
	"first-cosmos-sdk/testutil/nullify"
	"first-cosmos-sdk/x/firstcosmossdk"
	"first-cosmos-sdk/x/firstcosmossdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FirstcosmossdkKeeper(t)
	firstcosmossdk.InitGenesis(ctx, *k, genesisState)
	got := firstcosmossdk.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
