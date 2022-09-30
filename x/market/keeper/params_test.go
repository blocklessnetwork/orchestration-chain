package keeper_test

import (
	"testing"

	testkeeper "github.com/blocklessnetwork/orchestration-chain/testutil/keeper"
	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MarketKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
