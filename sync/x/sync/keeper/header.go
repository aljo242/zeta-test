package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync/x/sync/types"
)

// GetHeaderCount get the total number of header
func (k Keeper) GetHeaderCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HeaderCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHeaderCount set the total number of header
func (k Keeper) SetHeaderCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HeaderCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHeader appends a header in the store with a new id and update the count
func (k Keeper) AppendHeader(
	ctx sdk.Context,
	header types.Header,
) uint64 {
	// Create the header
	count := k.GetHeaderCount(ctx)

	// Set the ID of the appended value
	header.BlockID = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HeaderKey))
	appendedValue := k.cdc.MustMarshal(&header)
	store.Set(GetHeaderIDBytes(header.BlockID), appendedValue)

	// Update header count
	k.SetHeaderCount(ctx, count+1)

	return count
}

// SetHeader set a specific header in the store
func (k Keeper) SetHeader(ctx sdk.Context, header types.Header) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HeaderKey))
	b := k.cdc.MustMarshal(&header)
	store.Set(GetHeaderIDBytes(header.BlockID), b)
}

// GetHeader returns a header from its id
func (k Keeper) GetHeader(ctx sdk.Context, id uint64) (val types.Header, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HeaderKey))
	b := store.Get(GetHeaderIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHeader removes a header from the store
func (k Keeper) RemoveHeader(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HeaderKey))
	store.Delete(GetHeaderIDBytes(id))
}

// GetAllHeader returns all header
func (k Keeper) GetAllHeader(ctx sdk.Context) (list []types.Header) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HeaderKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Header
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHeaderIDBytes returns the byte representation of the ID
func GetHeaderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetHeaderIDFromBytes returns ID in uint64 format from a byte array
func GetHeaderIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
