package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/goccy/go-yaml"
	"github.com/pelletier/go-toml"
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
			Name:    "cache-dir",
			Aliases: []string{"c"},
			Usage:   "cache directory",
			Value:   config.CacheDir,
			EnvVars: []string{"PROTODEPS_CACHE_DIR", "CACHE_DIR"},
		},
		&cli.StringFlag{
			Name:    "output-dir",
			Aliases: []string{"o"},
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

func before(c *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		if _, err := os.Stat(".protodeps.json"); err == nil {
			b, err := ioutil.ReadFile(".protodeps.json")
			if err != nil {
				return err
			}
			if err := json.Unmarshal(b, &c); err != nil {
				return err
			}
		} else if _, err := os.Stat(".protodeps.yml"); err == nil {
			b, err := ioutil.ReadFile(".protodeps.yml")
			if err != nil {
				return err
			}
			if err := yaml.Unmarshal(b, &c); err != nil {
				return err
			}
		} else if _, err := os.Stat(".protodeps.yaml"); err == nil {
			b, err := ioutil.ReadFile(".protodeps.yaml")
			if err != nil {
				return err
			}
			if err := yaml.Unmarshal(b, &c); err != nil {
				return err
			}
		} else if _, err := os.Stat(".protodeps.toml"); err == nil {
			b, err := ioutil.ReadFile(".protodeps.toml")
			if err != nil {
				return err
			}
			if err := toml.Unmarshal(b, &c); err != nil {
				return err
			}
		}
		return nil
	}
}
