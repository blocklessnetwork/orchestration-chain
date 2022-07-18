package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/txlabs/blockless-chain/testutil/keeper"
	"github.com/txlabs/blockless-chain/x/market/keeper"
	"github.com/txlabs/blockless-chain/x/market/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MarketKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
