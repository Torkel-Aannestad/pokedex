package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) LocationExplore(area string) (LocationDetails, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area

	dat, exists := c.Cache.Get(url)
	if exists {
		locationDetails := LocationDetails{}
		err := json.Unmarshal(dat, &locationDetails)
		if err != nil {
			return LocationDetails{}, err
		}
		return locationDetails, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}
	res.Body.Close()

	locationDetails := LocationDetails{}
	err = json.Unmarshal(dat, &locationDetails)
	if err != nil {
		return LocationDetails{}, err
	}

	c.Cache.Add(url, dat)

	return locationDetails, nil
}
