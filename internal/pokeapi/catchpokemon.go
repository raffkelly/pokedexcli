package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func CatchData(c *pokecache.Cache, param string) (CatchPokemonData, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + param
	val, ok := c.Get(url)
	if ok {
		catchResponse := CatchPokemonData{}
		err := json.Unmarshal(val, &catchResponse)
		if err != nil {
			return CatchPokemonData{}, err
		}
		return catchResponse, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return CatchPokemonData{}, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return CatchPokemonData{}, fmt.Errorf("pokemon not found, status code: %v", res.StatusCode)
	}
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return CatchPokemonData{}, err
	}

	catchResponse := CatchPokemonData{}
	err = json.Unmarshal(dat, &catchResponse)
	if err != nil {
		return CatchPokemonData{}, err
	}
	c.Add(url, dat)
	return catchResponse, nil
}
