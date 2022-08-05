package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/txlabs/blockless-chain/x/market/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NodeRegistrationAll(c context.Context, req *types.QueryAllNodeRegistrationRequest) (*types.QueryAllNodeRegistrationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nodeRegistrations []types.NodeRegistration
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nodeRegistrationStore := prefix.NewStore(store, types.KeyPrefix(types.NodeRegistrationKeyPrefix))

	pageRes, err := query.Paginate(nodeRegistrationStore, req.Pagination, func(key []byte, value []byte) error {
		var nodeRegistration types.NodeRegistration
		if err := k.cdc.Unmarshal(value, &nodeRegistration); err != nil {
			return err
		}

		nodeRegistrations = append(nodeRegistrations, nodeRegistration)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNodeRegistrationResponse{NodeRegistration: nodeRegistrations, Pagination: pageRes}, nil
}

func (k Keeper) NodeRegistration(c context.Context, req *types.QueryGetNodeRegistrationRequest) (*types.QueryGetNodeRegistrationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNodeRegistration(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNodeRegistrationResponse{NodeRegistration: val}, nil
}
