package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "healthcheck/testutil/keeper"
	"healthcheck/x/healthcheck/keeper"
	"healthcheck/x/healthcheck/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMonitoredChainMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.HealthcheckKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateMonitoredChain{Creator: creator,
			ChainId: strconv.Itoa(i),
		}
		_, err := srv.CreateMonitoredChain(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetMonitoredChain(ctx,
			expected.ChainId,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestMonitoredChainMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateMonitoredChain
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateMonitoredChain{Creator: creator,
				ChainId: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateMonitoredChain{Creator: "B",
				ChainId: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateMonitoredChain{Creator: creator,
				ChainId: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.HealthcheckKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateMonitoredChain{Creator: creator,
				ChainId: strconv.Itoa(0),
			}
			_, err := srv.CreateMonitoredChain(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateMonitoredChain(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetMonitoredChain(ctx,
					expected.ChainId,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestMonitoredChainMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteMonitoredChain
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteMonitoredChain{Creator: creator,
				ChainId: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteMonitoredChain{Creator: "B",
				ChainId: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteMonitoredChain{Creator: creator,
				ChainId: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.HealthcheckKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateMonitoredChain(wctx, &types.MsgCreateMonitoredChain{Creator: creator,
				ChainId: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteMonitoredChain(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetMonitoredChain(ctx,
					tc.request.ChainId,
				)
				require.False(t, found)
			}
		})
	}
}
