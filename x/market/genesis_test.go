package market_test

import (
	"testing"

	keepertest "github.com/blocklessnetwork/orchestration-chain/testutil/keeper"
	"github.com/blocklessnetwork/orchestration-chain/testutil/nullify"
	"github.com/blocklessnetwork/orchestration-chain/x/market"
	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
	"github.com/stretchr/testify/require"
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
		ClaimedOrderList: []types.ClaimedOrder{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		NodeRegistrationList: []types.NodeRegistration{
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
	require.ElementsMatch(t, genesisState.ClaimedOrderList, got.ClaimedOrderList)
	require.ElementsMatch(t, genesisState.NodeRegistrationList, got.NodeRegistrationList)
	// this line is used by starport scaffolding # genesis/test/assert
}
