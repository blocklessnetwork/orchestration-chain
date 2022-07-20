package market_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/txlabs/blockless-chain/testutil/keeper"
	"github.com/txlabs/blockless-chain/testutil/nullify"
	"github.com/txlabs/blockless-chain/x/market"
	"github.com/txlabs/blockless-chain/x/market/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		OrderList: []types.Order{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CompletedOrderList: []types.CompletedOrder{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarketKeeper(t)
	market.InitGenesis(ctx, *k, genesisState)
	got := market.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.OrderList, got.OrderList)
	require.ElementsMatch(t, genesisState.CompletedOrderList, got.CompletedOrderList)
	// this line is used by starport scaffolding # genesis/test/assert
}
