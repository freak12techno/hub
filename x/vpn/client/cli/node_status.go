package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagVPNId  = "vpn-node-id"
	flagStatus = "status"
)

func ChangeNodeStatusCommand(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-node-status",
		Short: "Update VPN node status",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(authCli.GetAccountDecoder(cdc))

			vpnId := viper.GetString(flagVPNId)
			status := viper.GetBool(flagStatus)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			from, err := cliCtx.GetFromAddress()

			if err != nil {
				return err
			}

			vpnIdBytes, err := csdkTypes.AccAddressFromBech32(vpnId)

			if err != nil {
				return err
			}

			msg := vpn.NewNodeStatusMsg(from, vpnIdBytes.String(), status)

			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagVPNId, "", "VPN node ID")
	cmd.Flags().Bool(flagStatus, false, "Node status")

	return cmd

}
