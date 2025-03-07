package main

import (
	"time"

	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func main() {
	c := pokecache.NewCache(time.Second * 5)
	startRepl(c)
}
