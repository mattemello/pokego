package main

import (
	"github.com/matteomello/commands"
)

type cliCommand struct {
	command     string
	description string
	callBack    func() error
}

var commandMap = map[string]cliCommand{

	"help": {

		command:     "help",
		description: "display a help message",
		callBack:    commands.callHelp,
	},

	"exit": {

		command:     "exit",
		description: "exit the program",
		callBack:    commands.callExit,
	},
}

func main() {

}
