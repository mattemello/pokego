package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome in the pokedex")
	for {
		fmt.Print("pokedex > ")
		reader.Scan()

		word := cleanInput(reader.Text())

		if len(word) == 0 {
			continue
		}

		commandName := word[0]

		command, exist := getCommands()[commandName]
		if exist {

			err := command.callBack()

			if err != nil {
				fmt.Println(err)
				continue
			}

		} else {

			fmt.Println("command not found")
			continue

		}
	}
}

func cleanInput(text string) []string {

	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	command     string
	description string
	callBack    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{

		"help": {

			command:     "help",
			description: "display a help message",
			callBack:    callHelp,
		},

		"exit": {

			command:     "exit",
			description: "exit the program",
			callBack:    callExit,
		},
		"map": {
			command:     "map",
			description: "display the next 20 location",
			callBack:    callMap,
		},
		"mapb": {
			command:     "mapb",
			description: "display the previus 20 location",
			callBack:    callMapB,
		},
	}
}
