package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"healthcheck/testutil/sample"
)

func TestMsgCreateMonitoredChains_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateMonitoredChains
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateMonitoredChains{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateMonitoredChains{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateMonitoredChains_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateMonitoredChains
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateMonitoredChains{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateMonitoredChains{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteMonitoredChains_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteMonitoredChains
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteMonitoredChains{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteMonitoredChains{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
