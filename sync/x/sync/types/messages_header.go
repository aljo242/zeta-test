package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

const (
	TypeMsgCreateHeader = "create_header"
	TypeMsgUpdateHeader = "update_header"
	TypeMsgDeleteHeader = "delete_header"
)

var _ sdk.Msg = &MsgCreateHeader{}

func NewMsgCreateHeader(admin string, parentHash, uncleHash, rootHash, txHash, receiptHash, hash []byte) *MsgCreateHeader {
	return &MsgCreateHeader{
		Admin:       admin,
		ParentHash:  parentHash,
		UncleHash:   uncleHash,
		RootHash:    rootHash,
		TxHash:      txHash,
		ReceiptHash: receiptHash,
		Hash:        hash,
	}
}

func (msg *MsgCreateHeader) Route() string {
	return RouterKey
}

func (msg *MsgCreateHeader) Type() string {
	return TypeMsgCreateHeader
}

func (msg *MsgCreateHeader) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgCreateHeader) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHeader) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	const (
		lenHash = ethcommon.HashLength
	)

	if len(msg.ParentHash) != lenHash {
		return sdkerrors.Wrapf(ErrInvalidHash, "invalid length (%s)", err)
	}
	if len(msg.UncleHash) != lenHash {
		return sdkerrors.Wrapf(ErrInvalidHash, "invalid length (%s)", err)
	}
	if len(msg.RootHash) != lenHash {
		return sdkerrors.Wrapf(ErrInvalidHash, "invalid length (%s)", err)
	}
	if len(msg.TxHash) != lenHash {
		return sdkerrors.Wrapf(ErrInvalidHash, "invalid length (%s)", err)
	}
	if len(msg.ReceiptHash) != lenHash {
		return sdkerrors.Wrapf(ErrInvalidHash, "invalid length (%s)", err)
	}
	if len(msg.Hash) != lenHash {
		return sdkerrors.Wrapf(ErrInvalidHash, "invalid length (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateHeader{}

func NewMsgUpdateHeader(admin string, id uint64) *MsgUpdateHeader {
	return &MsgUpdateHeader{
		BlockID: id,
		Admin:   admin,
	}
}

func (msg *MsgUpdateHeader) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHeader) Type() string {
	return TypeMsgUpdateHeader
}

func (msg *MsgUpdateHeader) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgUpdateHeader) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHeader) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHeader{}

func NewMsgDeleteHeader(admin string, id uint64) *MsgDeleteHeader {
	return &MsgDeleteHeader{
		BlockID: id,
		Admin:   admin,
	}
}
func (msg *MsgDeleteHeader) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHeader) Type() string {
	return TypeMsgDeleteHeader
}

func (msg *MsgDeleteHeader) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgDeleteHeader) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHeader) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	return nil
}
