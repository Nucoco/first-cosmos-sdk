package keeper_test

import (
	"context"
	"testing"

	keepertest "first-cosmos-sdk/testutil/keeper"
	"first-cosmos-sdk/x/firstcosmossdk/keeper"
	"first-cosmos-sdk/x/firstcosmossdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FirstcosmossdkKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
