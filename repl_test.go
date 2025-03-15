package main

import  "testing"

func TestCleanInput(t *testing.T){
    cases := []struct{
    input string
    expected []string
    }{
        {
            input: "  hello world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "Charmander Bulbasaur PIKACHU",
            expected: []string{"charmander","bulbasaur","pikachu"},
        },
        {
            input: "  Charmander   Bulbasaur PIKACHU",
            expected: []string{"charmander","bulbasaur","pikachu"},
        },
}

    for _, c := range cases{
        actual := cleanInput((c.input))
        if len(actual) != len(c.expected){
            t.Errorf("length of the strings do not match %d != %d ", len(actual), len(c.expected))
        }
        for i := range actual{
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord{
                t.Errorf("words do not match %s != %s", word, expectedWord)
            }
        }
    }
}


