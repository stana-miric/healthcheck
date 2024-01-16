package cli

import (
	"context"

	"healthcheck/x/healthcheck/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListMonitoredChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-monitored-chains",
		Short: "list all MonitoredChains",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMonitoredChainRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MonitoredChainAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowMonitoredChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-monitored-chains [chain-id]",
		Short: "shows a MonitoredChain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChainId := args[0]

			params := &types.QueryGetMonitoredChainRequest{
				ChainId: argChainId,
			}

			res, err := queryClient.MonitoredChain(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
