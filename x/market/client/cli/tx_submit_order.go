package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/txlabs/blockless-chain/x/market/types"
)

var _ = strconv.Itoa(0)

func CmdSubmitOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-order [function-id] [fuel]",
		Short: "Broadcast message submit-order",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFunctionId := args[0]
			argFuel := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitOrder(
				clientCtx.GetFromAddress().String(),
				argFunctionId,
				argFuel,
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
