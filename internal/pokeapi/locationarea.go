package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func LocationData(config *Config, c *pokecache.Cache) (LocationAreaData, error) {
	val, ok := c.Get(*config.Next)
	if ok {
		locationResponse := LocationAreaData{}
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return LocationAreaData{}, err
		}
		return locationResponse, nil
	}

	res, err := http.Get(*config.Next)
	if err != nil {
		return LocationAreaData{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaData{}, err
	}

	locationResponse := LocationAreaData{}
	err = json.Unmarshal(dat, &locationResponse)
	if err != nil {
		return LocationAreaData{}, err
	}
	c.Add(*config.Next, dat)
	return locationResponse, nil
}
