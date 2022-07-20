package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/txlabs/blockless-chain/testutil/keeper"
	"github.com/txlabs/blockless-chain/testutil/nullify"
	"github.com/txlabs/blockless-chain/x/market/keeper"
	"github.com/txlabs/blockless-chain/x/market/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNClaimedOrder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ClaimedOrder {
	items := make([]types.ClaimedOrder, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetClaimedOrder(ctx, items[i])
	}
	return items
}

func TestClaimedOrderGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNClaimedOrder(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetClaimedOrder(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClaimedOrderRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNClaimedOrder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClaimedOrder(ctx,
			item.Index,
		)
		_, found := keeper.GetClaimedOrder(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestClaimedOrderGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNClaimedOrder(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClaimedOrder(ctx)),
	)
}
