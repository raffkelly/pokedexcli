package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/raffkelly/pokedexcli/internal/pokeapi"
	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func commandCatch(config *pokeapi.Config, c *pokecache.Cache, param string) error {
	if param == "" {
		return fmt.Errorf("must provide a pokemon to attempt to catch")
	}
	result, err := pokeapi.CatchData(c, param)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", param)

	//	Create a new random number generator with a custom seed (e.g., current time)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomChance := rng.Intn(700) + 1
	if randomChance < result.BaseExperience {
		fmt.Printf("%v escaped!\n", result.Name)
		return nil
	}
	fmt.Printf("%v was caught!\n", result.Name)
	fmt.Println("You may not inspect it with the inspect command.")
	pokedex[result.Name] = result
	return nil
}
