package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	genaccountsCli "github.com/cosmos/cosmos-sdk/x/genaccounts/client/cli"
	genutilCli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tm "github.com/tendermint/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub/app"
	"github.com/sentinel-official/hub/types"
)

const (
	flagInvarCheckPeriod = "invar-check-period"
)

var (
	invarCheckPeriod uint
)

func main() {
	cdc := app.MakeCodec()

	config := types.GetConfig()
	config.SetBech32PrefixForAccount(types.Bech32PrefixAccAddr, types.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(types.Bech32PrefixValAddr, types.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(types.Bech32PrefixConsAddr, types.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	cmd := &cobra.Command{
		Use:               "sentinel-hub-daemon",
		Short:             "Sentinel Hub Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	cmd.AddCommand(genutilCli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
	cmd.AddCommand(genutilCli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome))
	cmd.AddCommand(genutilCli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
		genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome))
	cmd.AddCommand(genutilCli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
	cmd.AddCommand(genaccountsCli.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	cmd.AddCommand(client.NewCompletionCmd(cmd, true))

	server.AddCommands(ctx, cdc, cmd, newApp, exportAppStateAndValidators)
	cmd.PersistentFlags().
		UintVar(&invarCheckPeriod, flagInvarCheckPeriod, 0, "Assert registered invariants every N blocks")

	executor := cli.PrepareBaseCmd(cmd, "SENTINEL_HUB", app.DefaultNodeHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db db.DB, tracer io.Writer) abci.Application {
	return app.NewApp(
		logger, db, tracer, true, invarCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportAppStateAndValidators(logger log.Logger, db db.DB, tracer io.Writer, height int64, zeroHeight bool,
	jailWhitelist []string) (json.RawMessage, []tm.GenesisValidator, error) {
	if height != -1 {
		hub := app.NewApp(logger, db, tracer, false, uint(1))
		if err := hub.LoadHeight(height); err != nil {
			return nil, nil, err
		}

		return hub.ExportAppStateAndValidators(zeroHeight, jailWhitelist)
	}

	hub := app.NewApp(logger, db, tracer, true, uint(1))
	return hub.ExportAppStateAndValidators(zeroHeight, jailWhitelist)
}
