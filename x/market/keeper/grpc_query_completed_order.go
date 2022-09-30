package keeper

import (
	"context"

	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CompletedOrderAll(c context.Context, req *types.QueryAllCompletedOrderRequest) (*types.QueryAllCompletedOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var completedOrders []types.CompletedOrder
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	completedOrderStore := prefix.NewStore(store, types.KeyPrefix(types.CompletedOrderKeyPrefix))

	pageRes, err := query.Paginate(completedOrderStore, req.Pagination, func(key []byte, value []byte) error {
		var completedOrder types.CompletedOrder
		if err := k.cdc.Unmarshal(value, &completedOrder); err != nil {
			return err
		}

		completedOrders = append(completedOrders, completedOrder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCompletedOrderResponse{CompletedOrder: completedOrders, Pagination: pageRes}, nil
}

func (k Keeper) CompletedOrder(c context.Context, req *types.QueryGetCompletedOrderRequest) (*types.QueryGetCompletedOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCompletedOrder(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCompletedOrderResponse{CompletedOrder: val}, nil
}
