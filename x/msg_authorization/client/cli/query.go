package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/msg_authorization/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	authorizationQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the msg authorization module",
		Long:                       "",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	authorizationQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryAuthorization(cdc),
	)...)

	return authorizationQueryCmd
}

// GetCmdQueryAuthorization implements the query authorizations command.
func GetCmdQueryAuthorization(cdc *codec.Codec) *cobra.Command {
	//TODO update description
	return &cobra.Command{
		Use:   "authorization",
		Args:  cobra.ExactArgs(3),
		Short: "query authorization for a granter-grantee pair",
		Long:  "query authorization for a granter-grantee pair",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.NewContext().WithCodec(cdc)

			granterAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			granteeAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msgAuthorized := args[2]

			res, _, err := clientCtx.QueryStore(types.GetActorAuthorizationKey(granteeAddr, granterAddr, msgAuthorized), types.QuerierRoute)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("no authorization found for given address pair ")
			}

			var grant types.AuthorizationGrant
			err = cdc.UnmarshalBinaryBare(res, &grant)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(grant)
		},
	}
}
