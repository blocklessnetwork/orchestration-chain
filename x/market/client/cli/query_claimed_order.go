package cli

import (
	"context"

	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListClaimedOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-claimed-order",
		Short: "list all claimedOrder",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllClaimedOrderRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ClaimedOrderAll(context.Background(), params)
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

func CmdShowClaimedOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-claimed-order [index]",
		Short: "shows a claimedOrder",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetClaimedOrderRequest{
				Index: argIndex,
			}

			res, err := queryClient.ClaimedOrder(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
