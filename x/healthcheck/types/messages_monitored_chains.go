package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMonitoredChains = "create_monitored_chains"
	TypeMsgUpdateMonitoredChains = "update_monitored_chains"
	TypeMsgDeleteMonitoredChains = "delete_monitored_chains"
)

var _ sdk.Msg = &MsgCreateMonitoredChains{}

func NewMsgCreateMonitoredChains(
	creator string,
	chainId string,
	connectionId string,

) *MsgCreateMonitoredChains {
	return &MsgCreateMonitoredChains{
		Creator:      creator,
		ChainId:      chainId,
		ConnectionId: connectionId,
	}
}

func (msg *MsgCreateMonitoredChains) Route() string {
	return RouterKey
}

func (msg *MsgCreateMonitoredChains) Type() string {
	return TypeMsgCreateMonitoredChains
}

func (msg *MsgCreateMonitoredChains) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMonitoredChains) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMonitoredChains) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMonitoredChains{}

func NewMsgUpdateMonitoredChains(
	creator string,
	chainId string,
	connectionId string,

) *MsgUpdateMonitoredChains {
	return &MsgUpdateMonitoredChains{
		Creator:      creator,
		ChainId:      chainId,
		ConnectionId: connectionId,
	}
}

func (msg *MsgUpdateMonitoredChains) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMonitoredChains) Type() string {
	return TypeMsgUpdateMonitoredChains
}

func (msg *MsgUpdateMonitoredChains) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMonitoredChains) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMonitoredChains) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMonitoredChains{}

func NewMsgDeleteMonitoredChains(
	creator string,
	chainId string,

) *MsgDeleteMonitoredChains {
	return &MsgDeleteMonitoredChains{
		Creator: creator,
		ChainId: chainId,
	}
}
func (msg *MsgDeleteMonitoredChains) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMonitoredChains) Type() string {
	return TypeMsgDeleteMonitoredChains
}

func (msg *MsgDeleteMonitoredChains) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMonitoredChains) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMonitoredChains) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
