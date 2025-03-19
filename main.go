package main

import (
	"time"
	"github.com/AnuragNegii/Pokedex/internal/pokecache"
    "fmt"
    "net/http"
    "io"
)

var cache *pokecache.Cache

func main(){
    config := &Config{}
    cache = pokecache.NewCache(5 * time.Minute)
    startRepl(config)
}

func getPokeAPIData(url string) ([]byte, error) {
    // Check if we have this URL cached
    if cachedData, found := cache.Get(url); found {
        fmt.Println("Using cached data for:", url)
        return cachedData, nil
    }

    // If not in cache, make the HTTP request
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    // Add the response to the cache
    cache.Add(url, data)
    
    return data, nil
}
