package keeper

import (
	"context"
	"fmt"

	"cifer/x/cifer/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMintdata(goCtx context.Context, msg *types.MsgCreateMintdata) (*types.MsgCreateMintdataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var mintdata = types.Mintdata{
		Creator:        msg.Creator,
		CreatorAddress: msg.CreatorAddress,
		Metadata:       msg.Metadata,
		Amount:         msg.Amount,
	}

	id := k.AppendMintdata(
		ctx,
		mintdata,
	)

	return &types.MsgCreateMintdataResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMintdata(goCtx context.Context, msg *types.MsgUpdateMintdata) (*types.MsgUpdateMintdataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var mintdata = types.Mintdata{
		Creator:        msg.Creator,
		Id:             msg.Id,
		CreatorAddress: msg.CreatorAddress,
		Metadata:       msg.Metadata,
		Amount:         msg.Amount,
	}

	// Checks that the element exists
	val, found := k.GetMintdata(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetMintdata(ctx, mintdata)

	return &types.MsgUpdateMintdataResponse{}, nil
}

func (k msgServer) DeleteMintdata(goCtx context.Context, msg *types.MsgDeleteMintdata) (*types.MsgDeleteMintdataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMintdata(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMintdata(ctx, msg.Id)

	return &types.MsgDeleteMintdataResponse{}, nil
}
