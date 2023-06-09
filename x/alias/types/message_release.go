package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRelease = "release"

var _ sdk.Msg = &MsgRelease{}

func NewMsgRelease(creator string, name string) *MsgRelease {
	return &MsgRelease{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgRelease) Route() string {
	return RouterKey
}

func (msg *MsgRelease) Type() string {
	return TypeMsgRelease
}

func (msg *MsgRelease) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRelease) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRelease) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = ValidateAlias(msg.Name)
	if err != nil {
		return ErrInvalidAlias
	}

	return nil
}
