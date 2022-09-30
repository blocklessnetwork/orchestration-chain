package keeper

import (
	"context"
	"strconv"
	"time"

	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ClaimOrderWork(goCtx context.Context, msg *types.MsgClaimOrderWork) (*types.MsgClaimOrderWorkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, isFound := k.GetOrder(ctx, msg.OrderIndex)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no order with that id exists")
	}

	if order.State != "open" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "order is not open")
	}

	order.State = "claimed"
	k.SetOrder(ctx, order)

	claimedBy, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(err)
	}

	claimedOrder := types.ClaimedOrder{
		Index:         order.Index,
		ClaimedBy:     claimedBy.String(),
		Date:          time.Now().String(),
		ClaimedHeight: strconv.FormatInt(ctx.BlockHeight(), 10),
	}

	k.SetClaimedOrder(ctx, claimedOrder)

	return &types.MsgClaimOrderWorkResponse{}, nil
}
