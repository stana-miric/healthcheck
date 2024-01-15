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

func TestMonitoredChainsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.HealthcheckKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateMonitoredChains{Creator: creator,
			ChainId: strconv.Itoa(i),
		}
		_, err := srv.CreateMonitoredChains(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetMonitoredChains(ctx,
			expected.ChainId,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestMonitoredChainsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateMonitoredChains
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateMonitoredChains{Creator: creator,
				ChainId: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateMonitoredChains{Creator: "B",
				ChainId: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateMonitoredChains{Creator: creator,
				ChainId: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.HealthcheckKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateMonitoredChains{Creator: creator,
				ChainId: strconv.Itoa(0),
			}
			_, err := srv.CreateMonitoredChains(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateMonitoredChains(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetMonitoredChains(ctx,
					expected.ChainId,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestMonitoredChainsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteMonitoredChains
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteMonitoredChains{Creator: creator,
				ChainId: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteMonitoredChains{Creator: "B",
				ChainId: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteMonitoredChains{Creator: creator,
				ChainId: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.HealthcheckKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateMonitoredChains(wctx, &types.MsgCreateMonitoredChains{Creator: creator,
				ChainId: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteMonitoredChains(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetMonitoredChains(ctx,
					tc.request.ChainId,
				)
				require.False(t, found)
			}
		})
	}
}
