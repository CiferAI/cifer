package keeper

import (
	"context"

	"cifer/x/cifer/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MintdataAll(ctx context.Context, req *types.QueryAllMintdataRequest) (*types.QueryAllMintdataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mintdatas []types.Mintdata

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	mintdataStore := prefix.NewStore(store, types.KeyPrefix(types.MintdataKey))

	pageRes, err := query.Paginate(mintdataStore, req.Pagination, func(key []byte, value []byte) error {
		var mintdata types.Mintdata
		if err := k.cdc.Unmarshal(value, &mintdata); err != nil {
			return err
		}

		mintdatas = append(mintdatas, mintdata)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMintdataResponse{Mintdata: mintdatas, Pagination: pageRes}, nil
}

func (k Keeper) Mintdata(ctx context.Context, req *types.QueryGetMintdataRequest) (*types.QueryGetMintdataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	mintdata, found := k.GetMintdata(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMintdataResponse{Mintdata: mintdata}, nil
}
