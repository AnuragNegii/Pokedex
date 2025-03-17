package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *Config){
   scanner := bufio.NewScanner(os.Stdin) 
    for{
        fmt.Print("Pokedex > ") 
        scanner.Scan()
        text := cleanInput(scanner.Text())
        if len(text) == 0{
            continue
        }
             
        commandName := text[0]
        
        command, exists := getCommands()[commandName]
        if exists{
            err := command.callback(config)
            if err != nil {
                fmt.Println("Error:", err)
            }           
            continue
            }else{
            fmt.Println("Unknown command")
            continue
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

type cliCommand struct{
    name string
    description string
    callback func(*Config) error
}

func getCommands() map[string]cliCommand{
     return map[string]cliCommand{
        "exit":{
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
        "help":{
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
        "map":{
            name:"map",
            description: "location of next 20 areas",
            callback: commandMap,
        },
        "mapb":{
            name:"mapb",
            description: "location of previous 20 areas",
            callback: commandMapb,
        },
    }
}

