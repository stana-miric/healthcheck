package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/monitored module sentinel errors
var (
	ErrSample                   = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout     = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion           = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrInvalidChannelFlow       = sdkerrors.Register(ModuleName, 1502, "invalid message sent to channel end")
	ErrInvalidHandshakeMetadata = sdkerrors.Register(ModuleName, 1503, "invalid handshake metadata")
)
