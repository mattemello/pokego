package main

import (
    "fmt"
    "errors"
)

func explore(location string) error {

    if location == "nil" {
        return errors.New("you don't have pass argoments! you have to do it")
    }

    _, conained := locationPoke[location]

    if conained == false {
        return errors.New("location not founded")
    }

    for _, pokemon  := range locationPoke[location].PokemonEncounters {

        fmt.Printf("%s", pokemon.Pokemon.Name)
        fmt.Println()

    }

    return nil

}
