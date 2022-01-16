package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cagnotteApp/x/cagnotte/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group cagnotte queries under a subcommand
	cagnotteQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cagnotteQueryCmd.AddCommand(
		flags.GetCommands(

			GetCmdGetCagnotte("cagnotte", cdc),
			GetCmdListCagnotte("cagnotte", cdc),
		)...,
	)

	return cagnotteQueryCmd
}
