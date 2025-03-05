package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var words []string
	text = strings.ToLower(text)
	words = strings.Fields(text)
	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInputString := scanner.Text()
		userInputWords := cleanInput(userInputString)
		if len(userInputWords) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v\n", userInputWords[0])
	}
}
