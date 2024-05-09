package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Select a pokemon to catch")
	}
	if cfg.CurrentLocation == nil {
		return errors.New("pokemon not in your current area")
	}

	wildPokemon, err := cfg.PokeapiClient.Pokemon(args[0])
	if err != nil {
		return err
	}

	PokemonEncountered := false
	for _, p := range cfg.CurrentLocation.PokemonEncounters {
		if wildPokemon.Name == p.Pokemon.Name {
			PokemonEncountered = true
		}
	}
	if !PokemonEncountered {
		return errors.New("pokemon is not in the area")
	}

	fmt.Println("")
	fmt.Printf("Throwing a Pokeball at %v ...\n", wildPokemon.Name)

	time.Sleep(time.Second * 1)

	probability := 0.3
	catchChance := rand.Float64()
	if catchChance < probability {
		fmt.Printf("%v ran away\n", wildPokemon.Name)
		return nil
	}

	cfg.Pokedex.Add(wildPokemon.Name, wildPokemon)
	fmt.Printf("%v was caught!\n", wildPokemon.Name)

	return nil
}
