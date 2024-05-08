package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config) error {
	fmt.Println()
	fmt.Println("Locations:")

	locationList, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationList.Next
	config.prevLocationsURL = locationList.Previous

	for _, loc := range locationList.Results {
		fmt.Printf("%v\n", loc.Name)
	}
	return nil
}
func commandMapb(config *config) error {
	if config.prevLocationsURL == nil {
		return errors.New("your already on the first page")
	}

	fmt.Println()
	fmt.Println("Locations:")
	locationList, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationList.Next
	config.prevLocationsURL = locationList.Previous

	for _, loc := range locationList.Results {
		fmt.Printf("%v\n", loc.Name)
	}
	return nil
}
