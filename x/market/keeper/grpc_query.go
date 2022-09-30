package keeper

import (
	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
)

var _ types.QueryServer = Keeper{}
