package commands

import (
	"flag"

	"github.com/mitchellh/cli"
)

type BaseCommand struct {
	UI      cli.Ui
	flagSet *flag.FlagSet
}

