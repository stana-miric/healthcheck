package types

import (
	"testing"

	"healthcheck/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateMonitoredChain_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateMonitoredChain
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateMonitoredChain{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateMonitoredChain{
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

func TestMsgUpdateMonitoredChain_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateMonitoredChain
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateMonitoredChain{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateMonitoredChain{
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

func TestMsgDeleteMonitoredChain_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteMonitoredChain
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteMonitoredChain{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteMonitoredChain{
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
