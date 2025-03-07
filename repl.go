package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/raffkelly/pokedexcli/internal/pokeapi"
	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func startRepl(c *pokecache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)
	config := pokeapi.GetConfig()
	param := ""
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInputString := scanner.Text()
		userInputWords := cleanInput(userInputString)
		if len(userInputWords) == 0 {
			continue
		}
		command, exists := getCommands()[userInputWords[0]]
		if len(userInputWords) > 1 {
			param = userInputWords[1]
		}
		if exists {
			err := command.callback(config, c, param)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config, c *pokecache.Cache, param string) error
}

func cleanInput(text string) []string {
	var words []string
	text = strings.ToLower(text)
	words = strings.Fields(text)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display a list of pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
	}
}
