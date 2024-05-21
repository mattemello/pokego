package main

import (
	"errors"
	"fmt"
)

func callMapB(nul string) error {
	if timeCall == 1 {
		return errors.New("you are at the start, you can't go back")
	}

	numToArrive := 0

	for i := 19; i > numToArrive; i-- {
		fmt.Printf("%s", locationPoke[arrayTwenty[i]].Name)
		fmt.Println()
	}

	timeCall -= 19
	return nil
}
