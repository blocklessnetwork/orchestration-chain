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

func createNNodeRegistration(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NodeRegistration {
	items := make([]types.NodeRegistration, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNodeRegistration(ctx, items[i])
	}
	return items
}

func TestNodeRegistrationGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNNodeRegistration(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNodeRegistration(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNodeRegistrationRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNNodeRegistration(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNodeRegistration(ctx,
			item.Index,
		)
		_, found := keeper.GetNodeRegistration(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestNodeRegistrationGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNNodeRegistration(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNodeRegistration(ctx)),
	)
}
