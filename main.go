package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/urfave/cli/v2"

	"github.com/takumin/protodeps/commands/completion"
	"github.com/takumin/protodeps/config"
)

var (
	AppName  string = "protodeps"
	Usage    string = "protocol buffers dependencies manager"
	Version  string = "0.0.0"
	Revision string = "development"
)

func main() {
	config := config.NewConfig(
		config.OutputDir("proto"),
		config.CacheDir(path.Join(xdg.CacheHome, "protodeps")),
	)

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "config filepath",
			EnvVars: []string{"PROTODEPS_CONFIG", "CONFIG"},
		},
		&cli.StringFlag{
			Name:    "cache-dir",
			Aliases: []string{"cache"},
			Usage:   "cache directory",
			Value:   config.CacheDir,
			EnvVars: []string{"PROTODEPS_CACHE_DIR", "CACHE_DIR"},
		},
		&cli.StringFlag{
			Name:    "output-dir",
			Aliases: []string{"output"},
			Usage:   "output directory",
			Value:   config.OutputDir,
			EnvVars: []string{"PROTODEPS_OUTPUT_DIR", "OUTPUT_DIR"},
		},
	}

	cmds := []*cli.Command{}
	cmds = append(cmds, completion.NewCommands(config, flags)...)

	app := &cli.App{
		Name:                 AppName,
		Usage:                Usage,
		Version:              fmt.Sprintf("%s (%s)", Version, Revision),
		Flags:                flags,
		Commands:             cmds,
		Before:               before(config),
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func before(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		if err := cfg.LoadIfExist(
			ctx.String("config"),
			".protodeps.json",
			".protodeps.yaml",
			".protodeps.yml",
			".protodeps.toml",
		); err != nil {
			return err
		}

		if ctx.IsSet("cache-dir") {
			cfg.CacheDir = ctx.String("cache-dir")
		}

		if ctx.IsSet("output-dir") {
			cfg.OutputDir = ctx.String("output-dir")
		}

		return nil
	}
}
