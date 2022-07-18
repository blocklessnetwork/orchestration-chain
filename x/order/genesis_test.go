package order_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/txlabs/blockless-chain/testutil/keeper"
	"github.com/txlabs/blockless-chain/testutil/nullify"
	"github.com/txlabs/blockless-chain/x/order"
	"github.com/txlabs/blockless-chain/x/order/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OrderKeeper(t)
	order.InitGenesis(ctx, *k, genesisState)
	got := order.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
