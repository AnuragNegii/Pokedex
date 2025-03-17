package main

import (
	"encoding/json"
	"io"
	"net/http"
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

func commandMap(config *Config) error{
    url := "https://pokeapi.co/api/v2/location-area/"
    if config.Next != ""{
        url = config.Next
    }

    resp, err := http.Get(url)
    if err != nil{
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
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

func commandMapb(config *Config) error{
    if config.Previous == "" {
        fmt.Println("you are on first page.")
        return nil
    }

    url := config.Previous

    resp, err := http.Get(url)
    if err != nil{
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
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
