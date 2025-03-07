package main

import (
	"fmt"

	"github.com/raffkelly/pokedexcli/internal/pokeapi"
	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func commandPokedex(config *pokeapi.Config, c *pokecache.Cache, param string) error {
	fmt.Println("Your Pokedex:")
	for name := range pokedex {
		fmt.Printf(" - %v\n", name)
	}
	return nil
}
