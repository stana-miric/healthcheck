package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateMonitoredChains{}, "healthcheck/CreateMonitoredChains", nil)
	cdc.RegisterConcrete(&MsgUpdateMonitoredChains{}, "healthcheck/UpdateMonitoredChains", nil)
	cdc.RegisterConcrete(&MsgDeleteMonitoredChains{}, "healthcheck/DeleteMonitoredChains", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMonitoredChains{},
		&MsgUpdateMonitoredChains{},
		&MsgDeleteMonitoredChains{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
