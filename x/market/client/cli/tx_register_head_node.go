package cli

import (
	"strconv"

	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRegisterHeadNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-head-node [node-id] [node-port] [node-ip]",
		Short: "Broadcast message register-head-node",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNodeId := args[0]
			argNodePort := args[1]
			argNodeIp := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterHeadNode(
				clientCtx.GetFromAddress().String(),
				argNodeId,
				argNodePort,
				argNodeIp,
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
