package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/txlabs/blockless-chain/testutil/keeper"
	"github.com/txlabs/blockless-chain/x/order/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OrderKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
