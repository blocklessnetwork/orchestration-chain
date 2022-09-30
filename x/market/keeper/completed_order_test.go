package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/blocklessnetwork/orchestration-chain/testutil/keeper"
	"github.com/blocklessnetwork/orchestration-chain/testutil/nullify"
	"github.com/blocklessnetwork/orchestration-chain/x/market/keeper"
	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCompletedOrder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CompletedOrder {
	items := make([]types.CompletedOrder, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCompletedOrder(ctx, items[i])
	}
	return items
}

func TestCompletedOrderGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNCompletedOrder(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCompletedOrder(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCompletedOrderRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNCompletedOrder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCompletedOrder(ctx,
			item.Index,
		)
		_, found := keeper.GetCompletedOrder(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCompletedOrderGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNCompletedOrder(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCompletedOrder(ctx)),
	)
}
