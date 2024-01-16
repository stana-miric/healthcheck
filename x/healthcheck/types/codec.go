package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateMonitoredChain{}, "healthcheck/CreateMonitoredChain", nil)
	cdc.RegisterConcrete(&MsgUpdateMonitoredChain{}, "healthcheck/UpdateMonitoredChain", nil)
	cdc.RegisterConcrete(&MsgDeleteMonitoredChain{}, "healthcheck/DeleteMonitoredChain", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMonitoredChain{},
		&MsgUpdateMonitoredChain{},
		&MsgDeleteMonitoredChain{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
