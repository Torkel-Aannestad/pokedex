package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type respLocationList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) LocationList(pageurl *string) (respLocationList, error) {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	if pageurl != nil {
		url = *pageurl
	}
	//check cache and use if exists
	dat, exists := c.Cache.Get(url)
	if exists {
		fmt.Println("Cache USED")
		locationList := respLocationList{}
		err := json.Unmarshal(dat, &locationList)
		if err != nil {
			return respLocationList{}, err
		}

		return locationList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respLocationList{}, err
	}
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return respLocationList{}, err
	}

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return respLocationList{}, err
	}
	res.Body.Close()

	locationList := respLocationList{}
	err = json.Unmarshal(dat, &locationList)
	if err != nil {
		return respLocationList{}, err
	}
	fmt.Println("HTTP REQUEST SENT")
	c.Cache.Add(url, dat) //Update cache

	return locationList, nil
}
