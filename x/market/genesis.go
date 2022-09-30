package market

import (
	"github.com/blocklessnetwork/orchestration-chain/x/market/keeper"
	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the order
	for _, elem := range genState.OrderList {
		k.SetOrder(ctx, elem)
	}
	// Set all the completedOrder
	for _, elem := range genState.CompletedOrderList {
		k.SetCompletedOrder(ctx, elem)
	}
	// Set all the claimedOrder
	for _, elem := range genState.ClaimedOrderList {
		k.SetClaimedOrder(ctx, elem)
	}
	// Set all the nodeRegistration
	for _, elem := range genState.NodeRegistrationList {
		k.SetNodeRegistration(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.OrderList = k.GetAllOrder(ctx)
	genesis.CompletedOrderList = k.GetAllCompletedOrder(ctx)
	genesis.ClaimedOrderList = k.GetAllClaimedOrder(ctx)
	genesis.NodeRegistrationList = k.GetAllNodeRegistration(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
