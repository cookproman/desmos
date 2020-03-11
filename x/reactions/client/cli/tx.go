package cli

import (
	"bufio"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reactions/internal/types"
	"github.com/spf13/cobra"
)

// GetTxCmd set the tx commands
func GetTxCmd(_ string, cdc *codec.Codec) *cobra.Command {
	postsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Reactions transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	postsTxCmd.AddCommand(flags.PostCommands(
		GetCmdAddPostReaction(cdc),
		GetCmdRemovePostReaction(cdc),
	)...)

	return postsTxCmd
}

// GetCmdAddPostReaction is the CLI command for adding a like to a post
func GetCmdAddPostReaction(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-reaction [post-id] [value]",
		Short: "Adds a reaction to a post",
		Long: fmt.Sprintf(`
Add a reaction to the post having the given id with the specified value. 
The value can be anything as long as it is ASCII supported.

E.g. 
%s tx posts add-reaction 12 like --from jack
%s tx posts add-reaction 12 👍 --from jack
`, version.ClientName, version.ClientName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			postID, err := posts.ParsePostID(args[0])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
			}

			msg := types.NewMsgAddPostReaction(postID, args[1], cliCtx.GetFromAddress())
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdRemovePostReaction is the CLI command for removing a like from a post
func GetCmdRemovePostReaction(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-reaction [post-id] [value]",
		Short: "Removes an existing reaction from a post",
		Long: fmt.Sprintf(`
Removes the reaction having the given value from the post having the given id. 
The value can be anything as long as it is ASCII supported.

E.g. 
%s tx posts remove-reaction 12 like --from jack
%s tx posts remove-reaction 12 👍 --from jack
`, version.ClientName, version.ClientName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			postID, err := posts.ParsePostID(args[0])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, err.Error())
			}

			msg := types.NewMsgRemovePostReaction(postID, cliCtx.GetFromAddress(), args[1])
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
