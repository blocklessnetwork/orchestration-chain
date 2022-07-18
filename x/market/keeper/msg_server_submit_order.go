package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/txlabs/blockless-chain/x/market/types"
)

func (k msgServer) SubmitOrder(goCtx context.Context, msg *types.MsgSubmitOrder) (*types.MsgSubmitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitOrderResponse{}, nil
}
