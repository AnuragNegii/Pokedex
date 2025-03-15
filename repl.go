package main

import "strings"

func cleanInput(text string) []string{
    newtext := strings.Fields(text)
    for i, word := range newtext{
        newtext[i] = strings.ToLower(word)
    } 
    return newtext 
}
