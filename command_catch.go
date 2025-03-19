package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"slices"
)


type PokemonType struct{
    Name string `json:"name"`
    Base_experience int `json:"base_experience"`
}

func commandCatch(config *Config, name string, caught []string) error{
    url := "https://pokeapi.co/api/v2/pokemon/" + name
    if name == ""{
        fmt.Println("No Pokemon Entered")
        return nil
    }
    body, err := http.Get(url)
    if err != nil{
        return err
    }

    pokemonType := PokemonType{}
    data, err := io.ReadAll(body.Body)
    if err != nil{
        return err
    }
    err = json.Unmarshal(data, &pokemonType) 
    if(err != nil){
        return err
    }
    fmt.Printf("Throwing a Pokeball at %s...\n", pokemonType.Name)
    fmt.Printf("base_experience : %d\n", pokemonType.Base_experience)
    gotPokemon := CatchPokemon(pokemonType.Base_experience) 
    if gotPokemon{
        fmt.Printf("%s caught!\n", pokemonType.Name)
    }else{
        fmt.Printf("%s escaped!\n", pokemonType.Name)
    }
    
    if slices.Contains(caught, pokemonType.Name) {
        return nil
    }
    
    caught = append(caught, pokemonType.Name)
    return nil
}

func CatchPokemon(baseExperience int) bool{
    randomNumber := rand.Intn(100)
    //simple pokemon with less base baseExperience
    if (baseExperience < 100){
        if (randomNumber < 50){
            return true
        }else{
            return false
        }
    }
    //simple pokemon with baseExperience between 100 and 200
    if (baseExperience > 100 && baseExperience < 200){
        if randomNumber < 50{
            return false
        }else if randomNumber > 50 && randomNumber < 80{
            return true
        }else{
            return false
        }
    }

    //simple pokemon with baseExperience above 200
    if (baseExperience > 200){
        if randomNumber > 80 {
            return true
        }else{
            return false
        }
    }
    return false
}
