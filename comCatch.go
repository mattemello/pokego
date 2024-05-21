package main

import (
    "errors"
    "fmt"
    "io"
    "net/http"
    "encoding/json"
    "math/rand"
)

var pokedex = map[string]Pokemon{}

func catch(name string) error {

    if name == "nil" {
        return errors.New("you don't have pass argoments! you have to do it")
    }

    contained := false

    for _, area := range locationPoke {
        for _, poke := range area.PokemonEncounters {
            if poke.Pokemon.Name == name {
                contained = true
            }
        }
    }

    if !contained {
        return errors.New("pokemon not in the location explored")
    }

    res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name + "/")
    if err != nil {
        return errors.New("error in the get, the pokemon don't exist or the connection don't work")
    }

    fmt.Printf("Throwing the pokeball to %s", name)
    fmt.Println()

    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        return errors.New("errore in the read body")
    }

    holder := Pokemon{}

    err = json.Unmarshal(body, &holder)
    if err != nil {
        return errors.New("error in the unmarshal")
    }

    value :=  rand.Intn(holder.BaseExperience)

    if value < 20 {
        pokedex[holder.Name] = holder
        fmt.Printf("%s was cautch!", holder.Name)
        fmt.Println()
    } else {
        fmt.Printf("%s escaped :(", holder.Name)
        fmt.Println()
    }

    return nil

}
