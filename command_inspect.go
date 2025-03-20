package main

import (
	"encoding/json"
	"fmt"
)

type Stats struct{
    Name string `json:"name"`
    Height int `json:"height"`
    Weight int `json:"weight"`
    Abilities []Ability `json:"abilities"`
    Stats []BaseStat `json:"stats"`
}

type Ability struct{
    AbilityDetail struct{
        Name string `json:"name"`
    }`json:"ability"`
}
type BaseStat struct{
    BaseStat int `json:"base_stat"`
    Stat struct{
        Name string `json:"name"`    
    }`json:"stat"`
}

func commandInspect(config *Config, pokemonName string, caught *[]Pokemon) error{
    isCaught := false
    for _, poke := range *caught{
        if pokemonName == poke.Name{
            isCaught = true 
        }
    }
    if isCaught{
        url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
        if pokemonName == ""{ 
            fmt.Println("ente pokemonName")
            return nil
        }
        body, err := getPokeAPIData(url)

        if err != nil{
            return err
        }
                
        stats := Stats{}
        err = json.Unmarshal(body, &stats)
        if err != nil {
            return err
        }
        fmt.Printf("Name: %s\n", stats.Name)
        fmt.Printf("Height: %d\n", stats.Height)
        fmt.Printf("Weight: %d\n", stats.Weight)
        fmt.Println("Abilities:")
        for _, ability := range stats.Abilities{
            fmt.Printf("- %s\n", ability.AbilityDetail.Name)
        }
        for _, s := range stats.Stats{
            fmt.Printf("- %s: %d\n", s.Stat.Name, s.BaseStat)
        }
        return nil
    }
    fmt.Printf("%s is not caught yet...\n", pokemonName)
    return nil
}
