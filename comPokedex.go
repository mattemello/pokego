package main

import (
    "fmt"
    "errors"
)

func pokedexTool(null string) error {

    if len(pokedex) == 0 {
        return errors.New("pokedex empty")
    }
    
    fmt.Println("Your pokedex")
    for _, pokemon := range pokedex {
        fmt.Printf(" - %s\n", pokemon.Name)
    }

    return nil
}
