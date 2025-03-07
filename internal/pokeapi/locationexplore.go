package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/raffkelly/pokedexcli/internal/pokecache"
)

func ExploreData(c *pokecache.Cache, param string) (LocationExploreData, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + param
	val, ok := c.Get(url)
	if ok {
		exploreResponse := LocationExploreData{}
		err := json.Unmarshal(val, &exploreResponse)
		if err != nil {
			return LocationExploreData{}, err
		}
		return exploreResponse, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationExploreData{}, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return LocationExploreData{}, fmt.Errorf("location not found, status code: %v", res.StatusCode)
	}
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationExploreData{}, err
	}

	exploreResponse := LocationExploreData{}
	err = json.Unmarshal(dat, &exploreResponse)
	if err != nil {
		return LocationExploreData{}, err
	}
	c.Add(url, dat)
	return exploreResponse, nil
}
