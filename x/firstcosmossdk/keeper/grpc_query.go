package keeper

import (
	"first-cosmos-sdk/x/firstcosmossdk/types"
)

var _ types.QueryServer = Keeper{}
