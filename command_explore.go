package main

import (
	"fmt"

	"github.com/raffkelly/pokedexcli/internal/pokeapi"
	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func commandExplore(config *pokeapi.Config, c *pokecache.Cache, param string) error {
	if param == "" {
		return fmt.Errorf("must provide location")
	}
	result, err := pokeapi.ExploreData(c, param)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", param)
	if len(result.PokemonEncounters) == 0 {
		return fmt.Errorf("no pokemon in this area")
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range result.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
