package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAdmin = "admin"

var _ sdk.Msg = &MsgAdmin{}

func NewMsgAdmin(authority string) *MsgAdmin {
	return &MsgAdmin{
		Authority: authority,
	}
}

func (msg *MsgAdmin) Route() string {
	return RouterKey
}

func (msg *MsgAdmin) Type() string {
	return TypeMsgAdmin
}

func (msg *MsgAdmin) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
