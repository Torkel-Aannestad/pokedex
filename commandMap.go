package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config, args ...string) error {
	locationList, err := cfg.PokeapiClient.LocationList(cfg.NextUrl)
	if err != nil {
		return err
	}
	fmt.Println("Locations:")
	for _, loc := range locationList.Results {
		fmt.Printf("%v\n", loc.Name)
	}
	cfg.NextUrl = locationList.Next
	cfg.PreviousUrl = locationList.Previous
	return nil
}
func commandMapb(cfg *Config, args ...string) error {
	if cfg.PreviousUrl == nil {
		return errors.New("your are already on the first page")
	}
	locationList, err := cfg.PokeapiClient.LocationList(cfg.PreviousUrl)
	if err != nil {
		return err
	}
	fmt.Println("Locations:")
	for _, loc := range locationList.Results {
		fmt.Printf("%v\n", loc.Name)
	}
	cfg.NextUrl = locationList.Next
	cfg.PreviousUrl = locationList.Previous
	return nil
}
