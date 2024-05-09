package main

import (
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {

	pokedex, err := cfg.Pokedex.GetAll()
	if err != nil {
		return err
	}

	fmt.Println("Your Pokedex:")
	for _, p := range pokedex {
		fmt.Printf("-%v\n", p.Name)
	}

	return nil
}
