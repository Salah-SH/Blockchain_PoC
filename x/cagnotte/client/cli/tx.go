package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cagnotteApp/x/cagnotte/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cagnotteTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cagnotteTxCmd.AddCommand(flags.PostCommands(
		GetCmdGetCagnotte("cagnotte", cdc),
		GetCmdListCagnotte("cagnotte", cdc),
		GetCmdAddCagnotte(cdc),
		GetCmdCreateCagnotte(cdc),
		GetCmdCloseCagnotte(cdc),
		GetCmdConfirmPending(cdc),
	)...)

	return cagnotteTxCmd
}
