package main

import (
	"encoding/json"
	"fmt"
)

type PokemonA struct{
    PokemonEncounters []struct{
        Pokemon struct{
            Name string `json:"name"`
            URL string `json:"url"`
        }`json:"pokemon"`
    }`json:"pokemon_encounters"`
}

func commandExplore(config *Config, areaName string, caught *[]Pokemon) error{
    fmt.Printf("Exploring %s\n", areaName)
    if areaName == ""{
        fmt.Printf("no area mentioned\n")
        return nil
    }
    url := "https://pokeapi.co/api/v2/location-area/" + areaName
    
    body, err := getPokeAPIData(url)
    if err != nil{
        return err
    }
    pokemon := PokemonA{}
    err = json.Unmarshal(body, &pokemon)
    if err != nil{
        return err
    }
    fmt.Println("Found Pokemon:")
    for _, encounter := range pokemon.PokemonEncounters{
        fmt.Printf("- %s\n", encounter.Pokemon.Name)
    }
    
    return nil
}
