package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMonitoredChain = "create_monitored_chains"
	TypeMsgUpdateMonitoredChain = "update_monitored_chains"
	TypeMsgDeleteMonitoredChain = "delete_monitored_chains"
)

var _ sdk.Msg = &MsgCreateMonitoredChain{}

func NewMsgCreateMonitoredChain(
	creator string,
	chainId string,
	connectionId string,

) *MsgCreateMonitoredChain {
	return &MsgCreateMonitoredChain{
		ChainId:      chainId,
		ConnectionId: connectionId,
	}
}

func (msg *MsgCreateMonitoredChain) Route() string {
	return RouterKey
}

func (msg *MsgCreateMonitoredChain) Type() string {
	return TypeMsgCreateMonitoredChain
}

func (msg *MsgCreateMonitoredChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMonitoredChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMonitoredChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMonitoredChain{}

func NewMsgUpdateMonitoredChain(
	creator string,
	chainId string,
	connectionId string,

) *MsgUpdateMonitoredChain {
	return &MsgUpdateMonitoredChain{
		Creator:      creator,
		ChainId:      chainId,
		ConnectionId: connectionId,
	}
}

func (msg *MsgUpdateMonitoredChain) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMonitoredChain) Type() string {
	return TypeMsgUpdateMonitoredChain
}

func (msg *MsgUpdateMonitoredChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMonitoredChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMonitoredChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMonitoredChain{}

func NewMsgDeleteMonitoredChain(
	creator string,
	chainId string,

) *MsgDeleteMonitoredChain {
	return &MsgDeleteMonitoredChain{
		Creator: creator,
		ChainId: chainId,
	}
}
func (msg *MsgDeleteMonitoredChain) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMonitoredChain) Type() string {
	return TypeMsgDeleteMonitoredChain
}

func (msg *MsgDeleteMonitoredChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMonitoredChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMonitoredChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
