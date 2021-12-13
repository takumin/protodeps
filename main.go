package main

import (
	"fmt"
	"log"
	"os"

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
	config := config.NewConfig()

	flags := []cli.Flag{}

	cmds := []*cli.Command{}
	cmds = append(cmds, completion.NewCommands(config, flags)...)

	app := &cli.App{
		Name:                 AppName,
		Usage:                Usage,
		Version:              fmt.Sprintf("%s (%s)", Version, Revision),
		Flags:                flags,
		Commands:             cmds,
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
