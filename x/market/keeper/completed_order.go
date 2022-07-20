package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/txlabs/blockless-chain/x/market/types"
)

// SetCompletedOrder set a specific completedOrder in the store from its index
func (k Keeper) SetCompletedOrder(ctx sdk.Context, completedOrder types.CompletedOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompletedOrderKeyPrefix))
	b := k.cdc.MustMarshal(&completedOrder)
	store.Set(types.CompletedOrderKey(
		completedOrder.Index,
	), b)
}

// GetCompletedOrder returns a completedOrder from its index
func (k Keeper) GetCompletedOrder(
	ctx sdk.Context,
	index string,

) (val types.CompletedOrder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompletedOrderKeyPrefix))

	b := store.Get(types.CompletedOrderKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCompletedOrder removes a completedOrder from the store
func (k Keeper) RemoveCompletedOrder(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompletedOrderKeyPrefix))
	store.Delete(types.CompletedOrderKey(
		index,
	))
}

// GetAllCompletedOrder returns all completedOrder
func (k Keeper) GetAllCompletedOrder(ctx sdk.Context) (list []types.CompletedOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompletedOrderKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CompletedOrder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
