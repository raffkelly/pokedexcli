package main

import (
	"time"

	"github.com/raffkelly/pokedexcli/internal/pokeapi"
	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

var pokedex map[string]pokeapi.CatchPokemonData

func main() {
	pokedex = make(map[string]pokeapi.CatchPokemonData)
	c := pokecache.NewCache(time.Second * 5)
	startRepl(c)
}
