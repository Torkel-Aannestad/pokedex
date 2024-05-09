package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonDetails, err := cfg.Pokedex.Get(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Name: %v\n", pokemonDetails.Name)
	fmt.Printf("Height: %v\n", pokemonDetails.Height)
	fmt.Printf("Weight: %v\n", pokemonDetails.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemonDetails.Stats {
		fmt.Printf("-%v: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemonDetails.Types {
		fmt.Printf("-%v\n", v.Type.Name)
	}

	return nil
}
