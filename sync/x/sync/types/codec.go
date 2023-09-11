package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateHeader{}, "sync/CreateHeader", nil)
	cdc.RegisterConcrete(&MsgUpdateHeader{}, "sync/UpdateHeader", nil)
	cdc.RegisterConcrete(&MsgDeleteHeader{}, "sync/DeleteHeader", nil)
	cdc.RegisterConcrete(&MsgAdmin{}, "sync/Admin", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHeader{},
		&MsgUpdateHeader{},
		&MsgDeleteHeader{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAdmin{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
