package main

import (
	"fmt"

	"github.com/raffkelly/pokedexcli/internal/pokeapi"
	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func commandMapf(config *pokeapi.Config, c *pokecache.Cache, param string) error {
	result, err := pokeapi.LocationData(config, c)
	if err != nil {
		return err
	}
	config.Next = result.Next
	config.Prev = result.Previous
	for _, location := range result.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(config *pokeapi.Config, c *pokecache.Cache, param string) error {
	if config.Prev == nil {
		return fmt.Errorf("already on first page of results")
	}
	tempconfig := *config
	tempconfig.Next = tempconfig.Prev
	result, err := pokeapi.LocationData(&tempconfig, c)
	if err != nil {
		return err
	}
	config.Next = result.Next
	config.Prev = result.Previous
	for _, location := range result.Results {
		fmt.Println(location.Name)
	}
	return nil
}
