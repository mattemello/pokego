package main

import "fmt"

func callHelp(nul string) error {

	fmt.Println()
	fmt.Println("the commands are: ")
	fmt.Println()

	for _, value := range getCommands() {
		fmt.Println("name: ", value.command)
		fmt.Println("description: ", value.description)
		fmt.Println()
	}

	return nil
}
