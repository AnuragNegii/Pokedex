package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(){
    scanner := bufio.NewScanner(os.Stdin) 
    for{
        fmt.Print("Pokedex > ") 
        if scanner.Scan(){
            text := cleanInput(scanner.Text())
            if len(text) == 0{
                continue
            }
            fmt.Printf("Your command was: %s\n", text[0])
        }
    }
}
 
func cleanInput(text string) []string{
    newtext := strings.Fields(text)
    for i, word := range newtext{
        newtext[i] = strings.ToLower(word)
    } 
    return newtext 
}
