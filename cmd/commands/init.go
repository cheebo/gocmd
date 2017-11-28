package commands

import "strings"

type InitCommand struct {
	BaseCommand
	args       []string
}

func (cmd *InitCommand) Run(args []string) int {

	/**
		You command code here
	 */
	println("Init")

	return 0
}

func (cmd *InitCommand) Help() string {
	helpText := `
Usage: file init
 `

	return strings.TrimSpace(helpText)
}

func (cmd *InitCommand) Synopsis() string {
	return "Init command"
}

