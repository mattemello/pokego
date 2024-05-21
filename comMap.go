package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var timeCall = 1    
var locationPoke = map[string]pokeJson{}
var arrayTwenty [20]string

func callMap(nul string) error {

	for i := timeCall; i < 20+timeCall; i++ {
		res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + strconv.Itoa(i) + "/?limit=1&offset=1")
		if err != nil {
			fmt.Println("error in the get")
		}


		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Println("error in the get")
		}
        holder := pokeJson{}

        err = json.Unmarshal(body, &holder)
        if err != nil {
            fmt.Println("error in the get")
        }

        arrayTwenty[i - timeCall] = holder.Name
        locationPoke[holder.Name] = holder
        fmt.Printf("%s", locationPoke[holder.Name].Name)
		fmt.Println()
	}

	timeCall += 19
	return nil
}
