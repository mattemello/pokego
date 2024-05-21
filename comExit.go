package main

import (
	"errors"
	"os"
)

func callExit(nul string) error {
	os.Exit(0)

	return errors.New("program not terminate")
}
