package healthcheck

import (
	"fmt"
	"healthcheck/x/healthcheck/keeper"
	"healthcheck/x/healthcheck/types"
	commonTypes "healthcheck/x/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v6/modules/core/exported"
)

type IBCModule struct {
	keeper keeper.Keeper
}

func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {

	return "", sdkerrors.Wrap(types.ErrInvalidChannelFlow, "channel handshake must be initiated by monitored chain")
}

// OnChanOpenTry implements the IBCModule interface
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	handshakeMetadata string,
) (string, error) {

	// Require portID is the portID module is bound to
	boundPort := im.keeper.GetPort(ctx)
	if boundPort != portID {
		return "", sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	if counterparty.PortId != commonTypes.MonitoredPortID {
		return "", sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, commonTypes.MonitoredPortID)
	}

	var md commonTypes.HandshakeMetadata
	if err := types.ModuleCdc.UnmarshalJSON([]byte(handshakeMetadata), &md); err != nil {
		return "", sdkerrors.Wrapf(types.ErrInvalidHandshakeMetadata,
			"error unmarshalling ibc-ack metadata: \n%v; \nmetadata: %v", err, handshakeMetadata)
	}

	if md.Version != commonTypes.Version {
		return "", sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: got: %s, expected %s", md.Version, commonTypes.Version)
	}

	if err := im.keeper.InitializeMonitoredChain(ctx, connectionHops[0], md.TimeoutInterval, md.UpdateInterval); err != nil {
		return "", err
	}

	if err := im.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return "", err
	}

	return commonTypes.Version, nil
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	_,
	counterpartyVersion string,
) error {
	return sdkerrors.Wrap(types.ErrInvalidChannelFlow, "channel handshake must be initiated by monitored chain")
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {

	chainID, err := im.keeper.GetClientChainIdFromChannel(ctx, channelID)
	if err != nil {
		return err
	}

	// monitored chain id to channel id map is saved to be able to close the channel on chain healthcheck timeout
	im.keeper.SetChainToChannelMap(ctx, chainID, channelID)

	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Disallow user-initiated channel closing for channels
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	var ack ibcexported.Acknowledgement

	// this line is used by starport scaffolding # oracle/packet/module/recv

	var packetData commonTypes.HealthcheckUpdateData
	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &packetData); err != nil {
		return channeltypes.NewErrorAcknowledgement(sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error()))
	} else {
		ack = im.keeper.OnRecvHealthcheckPacket(ctx, modulePacket.DestinationChannel, packetData)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"HealthcheckPacket",
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", ack != nil)),
		),
	)

	return ack
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	return sdkerrors.Wrap(types.ErrInvalidChannelFlow, "cannot receive packet acknowledgement on a healthcheck channel end")
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return sdkerrors.Wrap(types.ErrInvalidChannelFlow, "cannot receive packet timeout on a healthcheck channel end")
}
