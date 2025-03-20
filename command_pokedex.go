package main

import "fmt"

func commandPokedex(config *Config, word string, pokemonCaught *[]Pokemon) error{
    fmt.Println("Your Pokedex:")
    for _, poke := range *pokemonCaught{
        fmt.Printf("- %s\n", poke.Name)
    }
    return nil
}
