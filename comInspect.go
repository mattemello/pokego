package main

import (
    "fmt"
    "errors"
)

func inspect(name string) error {

    if name == "nil" {
        return errors.New("you don't have pass argoments! you have to do it")
    }

    pokemonInf, isIn := pokedex[name]
    if !isIn {
        return errors.New("pokemon not cautch")
    }

    fmt.Printf("Name: %s \n", pokemonInf.Name) 
    fmt.Printf("Height: %d \n", pokemonInf.Height) 
    fmt.Printf("Weight: %d \n", pokemonInf.Weight) 
    fmt.Println("Stats: ")
    for _, stat := range pokemonInf.Stats {
        fmt.Printf("\t- %s: %d \n", stat.Stat.Name, stat.BaseStat) 

    }
    fmt.Println("Types: ")
    for _, typ := range pokemonInf.Types {
        fmt.Printf("\t- %s \n", typ.Type.Name) 
    }

    return nil

}
