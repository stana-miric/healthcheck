package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"healthcheck/testutil/network"
	"healthcheck/testutil/nullify"
	"healthcheck/x/healthcheck/client/cli"
	"healthcheck/x/healthcheck/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithMonitoredChainsObjects(t *testing.T, n int) (*network.Network, []types.MonitoredChains) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		monitoredChains := types.MonitoredChains{
			ChainId: strconv.Itoa(i),
		}
		nullify.Fill(&monitoredChains)
		state.MonitoredChainsList = append(state.MonitoredChainsList, monitoredChains)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.MonitoredChainsList
}

func TestShowMonitoredChains(t *testing.T) {
	net, objs := networkWithMonitoredChainsObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc      string
		idChainId string

		args []string
		err  error
		obj  types.MonitoredChains
	}{
		{
			desc:      "found",
			idChainId: objs[0].ChainId,

			args: common,
			obj:  objs[0],
		},
		{
			desc:      "not found",
			idChainId: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idChainId,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowMonitoredChains(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetMonitoredChainsResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.MonitoredChains)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.MonitoredChains),
				)
			}
		})
	}
}

func TestListMonitoredChains(t *testing.T) {
	net, objs := networkWithMonitoredChainsObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMonitoredChains(), args)
			require.NoError(t, err)
			var resp types.QueryAllMonitoredChainsResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.MonitoredChains), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.MonitoredChains),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMonitoredChains(), args)
			require.NoError(t, err)
			var resp types.QueryAllMonitoredChainsResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.MonitoredChains), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.MonitoredChains),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMonitoredChains(), args)
		require.NoError(t, err)
		var resp types.QueryAllMonitoredChainsResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.MonitoredChains),
		)
	})
}
