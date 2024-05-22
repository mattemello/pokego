package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func start() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome in the pokedex")
	fmt.Println()

	for {
		fmt.Print("pokedex > ")
		exec.Command("/bin/stty", "-F", "/dev/tty", "-icanon", "min", "1").Run()
		var b []byte = make([]byte, 1)
		os.Stdin.Read(b)

		if bytes.Compare(b, []byte(strconv.Itoa(24))) == 0 {
			fmt.Println("worked")
		}

		reader.Scan()

		word := cleanInput(reader.Text())

		fmt.Printf("word %s, b %s", word, b)

		if len(word) == 0 {
			continue
		}

		commandName := word[0]

		argument := "nil"

		if len(word) > 1 {
			argument = word[1]
		}

		command, exist := getCommands()[commandName]
		if exist {

			fmt.Println()
			err := command.callBack(argument)

			if err != nil {
				fmt.Println(err)
				fmt.Println()
				continue
			}
			fmt.Println()

		} else {

			fmt.Println()
			fmt.Println("command not found")
			fmt.Println()
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
	callBack    func(string) error
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
		"explore": {
			command:     "explore",
			description: "displal all the pokemon in a specific location",
			callBack:    explore,
		},
		"catch": {
			command:     "catch",
			description: "make you catch the pokemon in the location",
			callBack:    catch,
		},
		"inspect": {
			command:     "inspect",
			description: "display the statis of one pokemon that you have captured",
			callBack:    inspect,
		},
		"pokedex": {
			command:     "pokedex",
			description: "display all your pokemon",
			callBack:    pokedexTool,
		},
	}
}
