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

func (k msgServer) SubmitOrder(goCtx context.Context, msg *types.MsgSubmitOrder) (*types.MsgSubmitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	currentTime := time.Now()
	height := strconv.FormatInt(ctx.BlockHeight(), 10)
	order := types.Order{
		Index:      msg.Creator + "-" + msg.FunctionId + "-" + height,
		FunctionId: msg.FunctionId,
		Customer:   msg.Creator,
		Fuel:       msg.Fuel,
		Height:     height,
		Date:       currentTime.String(),
		State:      "open",
	}

	// try getting a order from the store using the solution hash as the key
	_, isFound := k.GetOrder(ctx, order.Index)

	// return an error if a order already exists in the store
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "order with that id already exists")
	}

	// get address of the order module account
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	// convert the message creator address from a string into sdk.AccAddress
	customer, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(err)
	}

	// convert tokens from string into sdk.Coins
	reward, err := sdk.ParseCoinsNormalized(order.Fuel)
	if err != nil {
		panic(err)
	}

	// send tokens from the scavenge creator to the module account
	sdkError := k.bankKeeper.SendCoins(ctx, customer, moduleAcct, reward)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetOrder(ctx, order)
	return &types.MsgSubmitOrderResponse{}, nil
}
