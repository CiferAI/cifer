package keeper

import (
	"context"
	"encoding/binary"

	"cifer/x/cifer/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetMintdataCount get the total number of mintdata
func (k Keeper) GetMintdataCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MintdataCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMintdataCount set the total number of mintdata
func (k Keeper) SetMintdataCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MintdataCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMintdata appends a mintdata in the store with a new id and update the count
func (k Keeper) AppendMintdata(
	ctx context.Context,
	mintdata types.Mintdata,
) uint64 {
	// Create the mintdata
	count := k.GetMintdataCount(ctx)

	// Set the ID of the appended value
	mintdata.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintdataKey))
	appendedValue := k.cdc.MustMarshal(&mintdata)
	store.Set(GetMintdataIDBytes(mintdata.Id), appendedValue)

	// Update mintdata count
	k.SetMintdataCount(ctx, count+1)

	return count
}

// SetMintdata set a specific mintdata in the store
func (k Keeper) SetMintdata(ctx context.Context, mintdata types.Mintdata) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintdataKey))
	b := k.cdc.MustMarshal(&mintdata)
	store.Set(GetMintdataIDBytes(mintdata.Id), b)
}

// GetMintdata returns a mintdata from its id
func (k Keeper) GetMintdata(ctx context.Context, id uint64) (val types.Mintdata, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintdataKey))
	b := store.Get(GetMintdataIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMintdata removes a mintdata from the store
func (k Keeper) RemoveMintdata(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintdataKey))
	store.Delete(GetMintdataIDBytes(id))
}

// GetAllMintdata returns all mintdata
func (k Keeper) GetAllMintdata(ctx context.Context) (list []types.Mintdata) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintdataKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mintdata
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMintdataIDBytes returns the byte representation of the ID
func GetMintdataIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.MintdataKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
