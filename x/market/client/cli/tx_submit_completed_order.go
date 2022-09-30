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

func CmdSubmitCompletedOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-completed-order [order-index]",
		Short: "Broadcast message submit-completed-order",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOrderIndex := args[0]
			argFuelUsed := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitCompletedOrder(
				clientCtx.GetFromAddress().String(),
				argOrderIndex,
				argFuelUsed,
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
