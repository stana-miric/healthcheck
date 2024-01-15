package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"healthcheck/x/healthcheck/types"
)

func CmdCreateMonitoredChains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-monitored-chains [chain-id] [connection-id]",
		Short: "Create a new MonitoredChains",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexChainId := args[0]

			// Get value arguments
			argConnectionId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMonitoredChains(
				clientCtx.GetFromAddress().String(),
				indexChainId,
				argConnectionId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateMonitoredChains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-monitored-chains [chain-id] [connection-id]",
		Short: "Update a MonitoredChains",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexChainId := args[0]

			// Get value arguments
			argConnectionId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMonitoredChains(
				clientCtx.GetFromAddress().String(),
				indexChainId,
				argConnectionId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteMonitoredChains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-monitored-chains [chain-id]",
		Short: "Delete a MonitoredChains",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexChainId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMonitoredChains(
				clientCtx.GetFromAddress().String(),
				indexChainId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
