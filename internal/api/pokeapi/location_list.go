package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bkohler93/pokedexcli/internal/app/pokedex_cli/constants"
)

type RespLocationsList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(url *string) (RespLocationsList, error) {
	var locationUrl string
	var locations RespLocationsList

	if url == nil {
		locationUrl = baseUrl + fmt.Sprintf("location-area/?limit=%d", constants.MapLocationLimit)
	} else {
		locationUrl = *url
	}

	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return locations, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", locationUrl, nil)
	if err != nil {
		return locations, nil
	}

	res, err := c.client.Do(req)
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&locations)

	data, err := json.Marshal(locations)
	if err != nil {
		return locations, err
	}

	c.cache.Add(locationUrl, data)

	return locations, nil
}
