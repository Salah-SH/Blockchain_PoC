package cli

import (
	"fmt"

	"github.com/cagnotteApp/x/cagnotte/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetCmdListCagnotte(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-cagnotte  [addressUser]",
		Short: "list all cagnotte of a specific user",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryListCagnotte, cliCtx.GetFromAddress().String()), nil)
			if err != nil {
				return nil
			}
			var out []types.Cagnotte
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetCagnotte(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-cagnotte [key]",
		Short: "Query a cagnotte by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetCagnotte, key), nil)
			if err != nil {

				return nil
			}

			var out types.Cagnotte
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdResolveName queries information about a name
func GetCmdResolveName(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "resolve [name]",
		Short: "resolve name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryResolveName, name), nil)
			if err != nil {
				return nil
			}

			var out types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
