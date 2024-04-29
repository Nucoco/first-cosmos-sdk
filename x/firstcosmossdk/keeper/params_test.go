package keeper_test

import (
	"testing"

	testkeeper "first-cosmos-sdk/testutil/keeper"
	"first-cosmos-sdk/x/firstcosmossdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FirstcosmossdkKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
