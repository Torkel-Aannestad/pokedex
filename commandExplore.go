package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location provided")
	}

	locationDetails, err := cfg.PokeapiClient.LocationExplore(args[0])
	if err != nil {
		return err
	}
	fmt.Println("")
	fmt.Printf("Exploring %v\n", locationDetails.Location.Name)
	fmt.Println("Pokemons found:")
	if len(locationDetails.PokemonEncounters) != 0 {
		for _, v := range locationDetails.PokemonEncounters {
			fmt.Printf("-%v\n", v.Pokemon.Name)
		}
	} else {
		fmt.Println("No pokemons found")
	}
	cfg.CurrentLocation = &locationDetails

	return nil
}
