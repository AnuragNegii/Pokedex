package main

import (
	"encoding/json"
    "fmt"
)

type LocationResponse struct{
    Count int `json:"count"`
    Next string `json:"next"`
    Previous string `json:"previous"`
    Results []struct{
        Name string `json:"name"`
        URL string `json:"url"`
    } `json:"results"` 
}

type Config struct{
    Next string
    Previous string
}

func commandMap(config *Config, areaName string) error{
    url := "https://pokeapi.co/api/v2/location-area/" 
    if config.Next != ""{
        url = config.Next
    }

    body, err := getPokeAPIData(url) 
    if err != nil{
        return err
    }
     
    locationResponse := LocationResponse{}
    err = json.Unmarshal(body, &locationResponse)
    if err != nil{
        return err
    }
    for _, loc := range locationResponse.Results{
        fmt.Println(loc.Name)
    }

    config.Next = locationResponse.Next
    config.Previous = locationResponse.Previous

    return nil
}

func commandMapb(config *Config, areaName string) error{
    if config.Previous == "" {
        fmt.Println("you are on first page.")
        return nil
    }

    url := config.Previous

    body, err := getPokeAPIData(url)
    if err != nil{
        return err
    }
     
    locationResponse := LocationResponse{}
    err = json.Unmarshal(body, &locationResponse)
    if err != nil{
        return err
    }
    for _, loc := range locationResponse.Results{
        fmt.Println(loc.Name)
    }

    config.Next = locationResponse.Next
    config.Previous = locationResponse.Previous

    return nil
}
