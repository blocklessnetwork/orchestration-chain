package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/txlabs/blockless-chain/x/market/types"
)

// SetNodeRegistration set a specific nodeRegistration in the store from its index
func (k Keeper) SetNodeRegistration(ctx sdk.Context, nodeRegistration types.NodeRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeRegistrationKeyPrefix))
	b := k.cdc.MustMarshal(&nodeRegistration)
	store.Set(types.NodeRegistrationKey(
		nodeRegistration.Index,
	), b)
}

// GetNodeRegistration returns a nodeRegistration from its index
func (k Keeper) GetNodeRegistration(
	ctx sdk.Context,
	index string,

) (val types.NodeRegistration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeRegistrationKeyPrefix))

	b := store.Get(types.NodeRegistrationKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNodeRegistration removes a nodeRegistration from the store
func (k Keeper) RemoveNodeRegistration(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeRegistrationKeyPrefix))
	store.Delete(types.NodeRegistrationKey(
		index,
	))
}

// GetAllNodeRegistration returns all nodeRegistration
func (k Keeper) GetAllNodeRegistration(ctx sdk.Context) (list []types.NodeRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeRegistrationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NodeRegistration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
