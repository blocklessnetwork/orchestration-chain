package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/txlabs/blockless-chain/x/market/types"
)

func (k msgServer) SubmitCompletedOrder(goCtx context.Context, msg *types.MsgSubmitCompletedOrder) (*types.MsgSubmitCompletedOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitCompletedOrderResponse{}, nil
}
