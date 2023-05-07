package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegister = "register"

var _ sdk.Msg = &MsgRegister{}

func NewMsgRegister(creator string, name string) *MsgRegister {
	return &MsgRegister{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgRegister) Route() string {
	return RouterKey
}

func (msg *MsgRegister) Type() string {
	return TypeMsgRegister
}

func (msg *MsgRegister) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegister) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegister) ValidateBasic() error {
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
