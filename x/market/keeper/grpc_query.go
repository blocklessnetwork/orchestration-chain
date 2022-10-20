package keeper

import (
	"orchestration-chain/x/market/types"
)

var _ types.QueryServer = Keeper{}
