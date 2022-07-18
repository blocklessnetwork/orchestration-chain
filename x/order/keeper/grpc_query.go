package keeper

import (
	"github.com/txlabs/blockless-chain/x/order/types"
)

var _ types.QueryServer = Keeper{}
