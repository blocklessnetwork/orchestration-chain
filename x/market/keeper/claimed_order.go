package keeper

import (
	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetClaimedOrder set a specific claimedOrder in the store from its index
func (k Keeper) SetClaimedOrder(ctx sdk.Context, claimedOrder types.ClaimedOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimedOrderKeyPrefix))
	b := k.cdc.MustMarshal(&claimedOrder)
	store.Set(types.ClaimedOrderKey(
		claimedOrder.Index,
	), b)
}

// GetClaimedOrder returns a claimedOrder from its index
func (k Keeper) GetClaimedOrder(
	ctx sdk.Context,
	index string,

) (val types.ClaimedOrder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimedOrderKeyPrefix))

	b := store.Get(types.ClaimedOrderKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClaimedOrder removes a claimedOrder from the store
func (k Keeper) RemoveClaimedOrder(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimedOrderKeyPrefix))
	store.Delete(types.ClaimedOrderKey(
		index,
	))
}

// GetAllClaimedOrder returns all claimedOrder
func (k Keeper) GetAllClaimedOrder(ctx sdk.Context) (list []types.ClaimedOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimedOrderKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClaimedOrder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
