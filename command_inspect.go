package main

import (
	"fmt"
	"strconv"

	"github.com/raffkelly/pokedexcli/internal/pokeapi"
	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func commandInspect(config *pokeapi.Config, c *pokecache.Cache, param string) error {
	if param == "" {
		return fmt.Errorf("must provide a pokemon to inpsect")
	}
	val, ok := pokedex[param]
	if ok {
		message := `Name: ` + val.Name + `
Height: ` + strconv.Itoa(val.Height) + `
Weight: ` + strconv.Itoa(val.Weight) + `
Stats:
  -hp: ` + strconv.Itoa(val.Stats[0].BaseStat) + `
  -attack: ` + strconv.Itoa(val.Stats[1].BaseStat) + `
  -defense: ` + strconv.Itoa(val.Stats[2].BaseStat) + `
  -special-attack: ` + strconv.Itoa(val.Stats[3].BaseStat) + `
  -special-defense: ` + strconv.Itoa(val.Stats[4].BaseStat) + `
  -speed: ` + strconv.Itoa(val.Stats[5].BaseStat) + `
Types:`
		fmt.Println(message)
		for _, poketype := range val.Types {
			fmt.Printf("  - %v\n", poketype.Type.Name)
		}
		return nil
	}
	return fmt.Errorf("you have not caught that pokemon")
}
