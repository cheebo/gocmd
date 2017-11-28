package commands

import (
	"os"

	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &InitCommand{
				BaseCommand: BaseCommand{
					UI: ui,
				},
			}, nil
		},
	}
}

