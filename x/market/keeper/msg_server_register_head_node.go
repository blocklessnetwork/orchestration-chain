package keeper

import (
	"context"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/txlabs/blockless-chain/x/market/types"
)

func (k msgServer) RegisterHeadNode(goCtx context.Context, msg *types.MsgRegisterHeadNode) (*types.MsgRegisterHeadNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	currentTime := time.Now()
	height := strconv.FormatInt(ctx.BlockHeight(), 10)

	registration := types.NodeRegistration{
		Index:     msg.Creator + "-" + msg.NodeId,
		NodeId:    msg.NodeId,
		NodePort:  msg.NodePort,
		NodeIp:    msg.NodeIp,
		NodeOwner: msg.Creator,
		Height:    height,
		Date:      currentTime.String(),
	}

	// try getting a order from the store using the solution hash as the key
	_, isFound := k.GetNodeRegistration(ctx, registration.Index)

	// return an error if a order already exists in the store
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "node registration with that id already exists, update instead")
	}

	k.SetNodeRegistration(ctx, registration)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterHeadNodeResponse{}, nil
}
