// Copyright 2014 The go-ethereum Authors
// Copyright 2018 Ngin project
// This file is part of Ngin.
//
// Ngin is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Ngin is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Ngin. If not, see <http://www.gnu.org/licenses/>.

// ngind is the official command-line client for Ngin.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/urfave/cli.v1"

	rtppf "github.com/ethereumproject/benchmark/rtprof"
	"github.com/NginProject/ngind/common"
	"github.com/NginProject/ngind/console"
	"github.com/NginProject/ngind/core"
	"github.com/NginProject/ngind/logger"
	"github.com/NginProject/ngind/metrics"
)

// Version is the application revision identifier. It can be set with the linker
// as in: go build -ldflags "-X main.Version="`git describe --tags`
var Version = "source"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	common.SetClientVersion(Version)
}

var versionCommand = cli.Command{
	Action: version,
	Name:   "version",
	Usage:  "Print ngind version numbers",
	Description: `
	The output of this command is supposed to be machine-readable.
			`,
}

var makeMlogDocCommand = cli.Command{
	Action: makeMLogDocumentation,
	Name:   "mdoc",
	Usage:  "Generate mlog documentation",
	Description: `
	Auto-generates documentation for all available mlog lines.
	Use -md switch to toggle markdown output (eg. for wiki).
	Arguments may be used to specify exclusive candidate components;
	so 'ngind mdoc -md discover' will generate markdown documentation only
	for the 'discover' component.
			`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "md",
			Usage: "Toggle markdown formatting",
		},
	},
}

func makeCLIApp() (app *cli.App) {
	app = cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Version = Version
	app.Usage = "the ngind command line interface"
	app.Action = ng
	app.HideVersion = true // we have a command to print the version

	app.Commands = []cli.Command{
		importCommand,
		exportCommand,
		dumpChainConfigCommand,
		upgradedbCommand,
		dumpCommand,
		rollbackCommand,
		recoverCommand,
		resetCommand,
		monitorCommand,
		accountCommand,
		walletCommand,
		consoleCommand,
		attachCommand,
		javascriptCommand,
		statusCommand,
		apiCommand,
		versionCommand,
		makeMlogDocCommand,
		buildAddrTxIndexCommand,
	}

	app.Flags = []cli.Flag{
		PprofFlag,
		PprofIntervalFlag,
		SputnikVMFlag,
		NodeNameFlag,
		UnlockedAccountFlag,
		PasswordFileFlag,
		AccountsIndexFlag,
		BootnodesFlag,
		DataDirFlag,
		DocRootFlag,
		KeyStoreDirFlag,
		ChainIdentityFlag,
		BlockchainVersionFlag,
		FastSyncFlag,
		AddrTxIndexFlag,
		AddrTxIndexAutoBuildFlag,
		CacheFlag,
		LightKDFFlag,
		JSpathFlag,
		ListenPortFlag,
		MaxPeersFlag,
		MaxPendingPeersFlag,
		CoinbaseFlag,
		GasPriceFlag,
		MinerThreadsFlag,
		MiningEnabledFlag,
		TargetGasLimitFlag,
		NATFlag,
		NatspecEnabledFlag,
		NoDiscoverFlag,
		NodeKeyFileFlag,
		NodeKeyHexFlag,
		RPCEnabledFlag,
		RPCListenAddrFlag,
		RPCPortFlag,
		RPCApiFlag,
		WSEnabledFlag,
		WSListenAddrFlag,
		WSPortFlag,
		WSApiFlag,
		WSAllowedOriginsFlag,
		IPCDisabledFlag,
		IPCApiFlag,
		IPCPathFlag,
		ExecFlag,
		PreloadJSFlag,
		WhisperEnabledFlag,
		DevModeFlag,
		TestNetFlag,
		NetworkIdFlag,
		RPCCORSDomainFlag,
		NeckbeardFlag,
		VerbosityFlag,
		DisplayFlag,
		DisplayFormatFlag,
		VModuleFlag,
		LogDirFlag,
		LogMaxSizeFlag,
		LogMinSizeFlag,
		LogMaxTotalSizeFlag,
		LogIntervalFlag,
		LogMaxAgeFlag,
		LogCompressFlag,
		LogStatusFlag,
		MLogFlag,
		MLogDirFlag,
		MLogComponentsFlag,
		BacktraceAtFlag,
		MetricsFlag,
		FakePoWFlag,
		SolcPathFlag,
		GpoMinGasPriceFlag,
		GpoMaxGasPriceFlag,
		GpoFullBlockRatioFlag,
		GpobaseStepDownFlag,
		GpobaseStepUpFlag,
		GpobaseCorrectionFactorFlag,
		ExtraDataFlag,
		Unused1,
	}

	app.Before = func(ctx *cli.Context) error {

		// It's a patch.
		// Don't know why urfave/cli isn't catching the unknown command on its own.
		if ctx.Args().Present() {
			commandExists := false
			for _, cmd := range app.Commands {
				if cmd.HasName(ctx.Args().First()) {
					commandExists = true
				}
			}
			if !commandExists {
				if e := cli.ShowCommandHelp(ctx, ctx.Args().First()); e != nil {
					return e
				}
			}
		}

		// Check for --exec set without console OR attach
		if ctx.IsSet(ExecFlag.Name) {
			// If no command is used, OR command is not one of the valid commands attach/console
			if cmdName := ctx.Args().First(); cmdName == "" || (cmdName != "console" && cmdName != "attach") {
				log.Printf("Error: --%v flag requires use of 'attach' OR 'console' command, command was: '%v'", ExecFlag.Name, cmdName)
				cli.ShowCommandHelp(ctx, consoleCommand.Name)
				cli.ShowCommandHelp(ctx, attachCommand.Name)
				os.Exit(1)
			}
		}

		if ctx.IsSet(SputnikVMFlag.Name) {
			if core.SputnikVMExists {
				core.UseSputnikVM = true
			} else {
				log.Fatal("This version of ngind wasn't built to include SputnikVM. To build with SputnikVM, use -tags=sputnikvm following the go build command.")
			}
		}

		if err := setupLogRotation(ctx); err != nil {
			return err
		}

		// Handle parsing and applying log verbosity, severities, and default configurations from context.
		if err := setupLogging(ctx); err != nil {
			return err
		}

		// Handle parsing and applying log rotation configs from context.
		if err := setupLogRotation(ctx); err != nil {
			return err
		}

		if s := ctx.String("metrics"); s != "" {
			go metrics.CollectToFile(s)
		}

		// (whilei): I use `log` instead of `glog` because git diff tells me:
		// > The output of this command is supposed to be machine-readable.
		gasLimit := ctx.GlobalString(aliasableName(TargetGasLimitFlag.Name, ctx))
		if _, ok := core.TargetGasLimit.SetString(gasLimit, 0); !ok {
			return fmt.Errorf("malformed %s flag value %q", aliasableName(TargetGasLimitFlag.Name, ctx), gasLimit)
		}

		// Set testnet chain by default for dev mode.
		if ctx.GlobalBool(aliasableName(DevModeFlag.Name, ctx)) {
			if !ctx.GlobalIsSet(aliasableName(ChainIdentityFlag.Name, ctx)) {
				if e := ctx.Set(aliasableName(ChainIdentityFlag.Name, ctx), "testnet"); e != nil {
					return fmt.Errorf("failed to set chain value: %v", e)
				}
			}
		}

		if port := ctx.GlobalInt(PprofFlag.Name); port != 0 {
			interval := 5 * time.Second
			if i := ctx.GlobalInt(PprofIntervalFlag.Name); i > 0 {
				interval = time.Duration(i) * time.Second
			}
			rtppf.Start(interval, port)
		}

		return nil
	}

	app.After = func(ctx *cli.Context) error {
		rtppf.Stop()
		logger.Flush()
		console.Stdin.Close() // Resets terminal mode.
		return nil
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Invalid command: %q. Please find `ngind` usage below. \n", command)
		cli.ShowAppHelp(c)
		os.Exit(3)
	}
	return app
}

func main() {
	app := makeCLIApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// ng is the main entry point into the system if no special subcommand is ran.
// It creates a default node based on the command line arguments and runs it in
// blocking mode, waiting for it to be shut down.
func ng(ctx *cli.Context) error {

	n := MakeSystemNode(Version, ctx)
	ethe := startNode(ctx, n)

	if ctx.GlobalString(LogStatusFlag.Name) != "off" {
		dispatchStatusLogs(ctx, ethe)
	}
	logLoggingConfiguration(ctx)

	n.Wait()

	return nil
}
