package completion

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/protodeps/commands/completion/bash"
	"github.com/takumin/protodeps/commands/completion/fish"
	"github.com/takumin/protodeps/commands/completion/powershell"
	"github.com/takumin/protodeps/commands/completion/zsh"
	"github.com/takumin/protodeps/config"
)

func NewCommands(c *config.Config, f []cli.Flag) []*cli.Command {
	cmds := []*cli.Command{}
	cmds = append(cmds, bash.NewCommands(c, f)...)
	cmds = append(cmds, fish.NewCommands(c, f)...)
	cmds = append(cmds, zsh.NewCommands(c, f)...)
	cmds = append(cmds, powershell.NewCommands(c, f)...)

	return []*cli.Command{
		{
			Name:        "completion",
			Usage:       "command completion",
			Subcommands: cmds,
			HideHelp:    true,
		},
	}
}
