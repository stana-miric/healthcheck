package cli

import (
	"healthcheck/x/healthcheck/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateMonitoredChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-monitored-chains [chain-id] [connection-id]",
		Short: "Create a new MonitoredChain",
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

			msg := types.NewMsgCreateMonitoredChain(
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

func CmdUpdateMonitoredChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-monitored-chains [chain-id] [connection-id]",
		Short: "Update a MonitoredChain",
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

			msg := types.NewMsgUpdateMonitoredChain(
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

func CmdDeleteMonitoredChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-monitored-chains [chain-id]",
		Short: "Delete a MonitoredChain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexChainId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMonitoredChain(
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
