package main

import (
	"fmt"
	"os"

	"github.com/raffkelly/pokedexcli/internal/pokeapi"
	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func commandExit(config *pokeapi.Config, c *pokecache.Cache, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Failed to exit.")
}
