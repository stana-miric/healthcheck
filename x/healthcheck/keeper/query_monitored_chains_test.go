package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "healthcheck/testutil/keeper"
	"healthcheck/testutil/nullify"
	"healthcheck/x/healthcheck/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMonitoredChainsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.HealthcheckKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMonitoredChains(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMonitoredChainsRequest
		response *types.QueryGetMonitoredChainsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMonitoredChainsRequest{
				ChainId: msgs[0].ChainId,
			},
			response: &types.QueryGetMonitoredChainsResponse{MonitoredChains: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMonitoredChainsRequest{
				ChainId: msgs[1].ChainId,
			},
			response: &types.QueryGetMonitoredChainsResponse{MonitoredChains: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMonitoredChainsRequest{
				ChainId: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MonitoredChains(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestMonitoredChainsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.HealthcheckKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMonitoredChains(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMonitoredChainsRequest {
		return &types.QueryAllMonitoredChainsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MonitoredChainsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MonitoredChains), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MonitoredChains),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MonitoredChainsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MonitoredChains), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MonitoredChains),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MonitoredChainsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.MonitoredChains),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MonitoredChainsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
