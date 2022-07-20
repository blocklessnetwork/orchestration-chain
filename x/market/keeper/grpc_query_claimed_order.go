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

func (k Keeper) ClaimedOrderAll(c context.Context, req *types.QueryAllClaimedOrderRequest) (*types.QueryAllClaimedOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var claimedOrders []types.ClaimedOrder
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	claimedOrderStore := prefix.NewStore(store, types.KeyPrefix(types.ClaimedOrderKeyPrefix))

	pageRes, err := query.Paginate(claimedOrderStore, req.Pagination, func(key []byte, value []byte) error {
		var claimedOrder types.ClaimedOrder
		if err := k.cdc.Unmarshal(value, &claimedOrder); err != nil {
			return err
		}

		claimedOrders = append(claimedOrders, claimedOrder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClaimedOrderResponse{ClaimedOrder: claimedOrders, Pagination: pageRes}, nil
}

func (k Keeper) ClaimedOrder(c context.Context, req *types.QueryGetClaimedOrderRequest) (*types.QueryGetClaimedOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetClaimedOrder(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetClaimedOrderResponse{ClaimedOrder: val}, nil
}
