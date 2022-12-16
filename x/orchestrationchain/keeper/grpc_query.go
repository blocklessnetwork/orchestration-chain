package keeper

import (
	"github.com/blocklessnetwork/orchestration-chain/x/orchestrationchain/types"
)

var _ types.QueryServer = Keeper{}
