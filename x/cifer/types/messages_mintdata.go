package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMintdata{}

func NewMsgCreateMintdata(creator string, creatorAddress string, metadata string, amount string) *MsgCreateMintdata {
	return &MsgCreateMintdata{
		Creator:        creator,
		CreatorAddress: creatorAddress,
		Metadata:       metadata,
		Amount:         amount,
	}
}

func (msg *MsgCreateMintdata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMintdata{}

func NewMsgUpdateMintdata(creator string, id uint64, creatorAddress string, metadata string, amount string) *MsgUpdateMintdata {
	return &MsgUpdateMintdata{
		Id:             id,
		Creator:        creator,
		CreatorAddress: creatorAddress,
		Metadata:       metadata,
		Amount:         amount,
	}
}

func (msg *MsgUpdateMintdata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMintdata{}

func NewMsgDeleteMintdata(creator string, id uint64) *MsgDeleteMintdata {
	return &MsgDeleteMintdata{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteMintdata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
