package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cagnotteApp/x/cagnotte/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetCmdAddCagnotte(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-cagnotte [name] [amount]",
		Short: "participate in a cagnotte",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			coins, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddCagnotte(argsName, coins, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdCreateCagnotte(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-cagnotte [name]",
		Short: "Create a cagnotte",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgCreateCagnotte(argsName, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdCloseCagnotte(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "close-cagnotte [name]",
		Short: "Close a cagnotte",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgCloseCagnotte(argsName, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
func GetCmdConfirmPending(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "confirm-tx [name] [amount] [user] [success]",
		Short: "Confirm transaction from admin",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			coins, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			user, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}
			succes, err := strconv.ParseBool(args[3])
			if err != nil {
				return err
			}
			msg := types.NewMsgConfirmPendingTx(argsName, coins, cliCtx.GetFromAddress(), succes, user)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
