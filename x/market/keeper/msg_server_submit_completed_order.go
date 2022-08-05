package keeper

import (
	"context"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
	"github.com/txlabs/blockless-chain/x/market/types"
)

func (k msgServer) SubmitCompletedOrder(goCtx context.Context, msg *types.MsgSubmitCompletedOrder) (*types.MsgSubmitCompletedOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	currentTime := time.Now()
	height := strconv.FormatInt(ctx.BlockHeight(), 10)

	completedOrder := types.CompletedOrder{
		Index:       msg.Creator + "-" + msg.OrderIndex,
		OrderIndex:  msg.OrderIndex,
		CompletedBy: msg.Creator,
		FuelUsed:    msg.FuelUsed,
		Height:      height,
		ResponseId:  msg.ResponseId,
		Date:        currentTime.String(),
	}

	// try getting a order from the store using the solution hash as the key
	_, isFound := k.GetCompletedOrder(ctx, completedOrder.Index)

	// return an error if a order already exists in the store
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "completed order with that id already exists")
	}

	// try getting a order from the store using the solution hash as the key
	order, isFound := k.GetOrder(ctx, completedOrder.OrderIndex)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no order with that orderIndex found")
	}

	if order.State != "claimed" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "order is not in claimed state")
	}

	claimedOrder, isFound := k.GetClaimedOrder(ctx, order.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "claimed order is not found")
	}

	if claimedOrder.ClaimedBy != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "claimed order is not claimed by the completed order creator")
	}

	// get address of the order module account
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	// convert the message creator address from a string into sdk.AccAddress
	completedBy, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(err)
	}

	// @todo ensure that the order fuel and fuel used are inline
	// @todo refund unused fuel back to the customer
	// convert tokens from string into sdk.Coins
	reward, err := sdk.ParseCoinsNormalized(order.Fuel)
	if err != nil {
		panic(err)
	}

	// send tokens from the scavenge creator to the module account
	sdkError := k.bankKeeper.SendCoins(ctx, moduleAcct, completedBy, reward)
	if sdkError != nil {
		return nil, sdkError
	}

	order.State = "completed"
	k.SetOrder(ctx, order)
	k.SetCompletedOrder(ctx, completedOrder)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitCompletedOrderResponse{}, nil
}
