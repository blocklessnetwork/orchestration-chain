package keeper

import (
	"github.com/txlabs/blockless-chain/x/market/types"
)

var _ types.QueryServer = Keeper{}
