package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (respLocationArea, error) {
	url := baseURL + "/location-area/?offset=0&limit=20"
	if pageUrl != nil {
		url = *pageUrl
	}

	//If in cache, use cache
	dat, exists := c.cache.Get(url)
	if exists {
		fmt.Printf("Cache used: %v\n", url)
		locations := respLocationArea{}
		errors := json.Unmarshal(dat, &locations)
		if errors != nil {
			return respLocationArea{}, errors
		}
		return locations, nil
	}

	fmt.Printf("New request: %v\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respLocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return respLocationArea{}, err
	}

	dat, err = io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return respLocationArea{}, err
	}

	locations := respLocationArea{}
	errors := json.Unmarshal(dat, &locations)
	if errors != nil {
		return respLocationArea{}, errors
	}

	c.cache.Add(url, dat) //Save to cache

	return locations, nil

}
