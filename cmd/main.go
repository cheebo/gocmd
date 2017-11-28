package main

import (
	"os"
	"fmt"
	"github.com/cheebo/basecmd/cmd/commands"
	"github.com/mitchellh/cli"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "--" {
			break
		}

		if arg == "-v" || arg == "--version" {
			args = []string{"version"}
			break
		}
	}

	var cmds []string
	for c := range commands.Commands {
		cmds = append(cmds, c)
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: commands.Commands,
		HelpFunc: cli.FilteredHelpFunc(cmds, cli.BasicHelpFunc("Service")),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
