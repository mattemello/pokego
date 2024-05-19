package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func callMapB() error {
	if timeCall == 1 {
		return errors.New("you are at the start, you can't go back")
	}

	numToArrive := 1

	if 20-timeCall > 0 {
		numToArrive = 20 - timeCall
	} else {
		numToArrive = timeCall - 20
	}

	for i := timeCall; i > numToArrive; i-- {
		res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + strconv.Itoa(i) + "/?limit=1&offset=1")
		if err != nil {
			fmt.Println("error in the get")
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Println("error in the get")
		}

		prova := pokeJson{}

		err = json.Unmarshal(body, &prova)
		if err != nil {
			fmt.Println("error in the get")
		}
		fmt.Printf("%s", prova.Name)
		fmt.Println()
	}

	timeCall -= 19
	return nil
}
