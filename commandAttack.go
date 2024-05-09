package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/torkelaannestad/pokedex/internal/pokeapi"
)

func commandAttack(cfg *Config, args ...string) error {
	if len(args) != 2 {
		return errors.New("select pokemon to catch and your selected pokemon")
	}
	if cfg.CurrentLocation == nil {
		return errors.New("pokemon not in your current area")
	}

	wildPokemon, err := cfg.PokeapiClient.Pokemon(args[0])
	if err != nil {
		return err
	}

	playerPokemon, err := cfg.Pokedex.Get(args[1])
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

	battleResult := battle(wildPokemon, playerPokemon)
	if !battleResult {
		fmt.Printf("%v won the battle\n", wildPokemon.Name)
		return nil
	}
	fmt.Println("")
	fmt.Printf("Throwing a Pokeball at %v ...\n", wildPokemon.Name)

	time.Sleep(time.Second * 1)

	cfg.Pokedex.Add(wildPokemon.Name, wildPokemon)
	fmt.Printf("%v was caught!\n", wildPokemon.Name)

	return nil
}

func battle(wildPokemon pokeapi.Pokemon, playerPokemon pokeapi.Pokemon) bool {
	playerHealth := playerPokemon.Stats[0].BaseStat
	playerAttack := playerPokemon.Stats[1].BaseStat
	playerDefence := playerPokemon.Stats[2].BaseStat
	playerAttackDamage := 0
	if playerPokemon.Stats[1].BaseStat-wildPokemon.Stats[2].BaseStat > 0 {
		playerAttackDamage = playerPokemon.Stats[1].BaseStat - wildPokemon.Stats[2].BaseStat
	}

	wildHealth := wildPokemon.Stats[0].BaseStat
	wildAttack := wildPokemon.Stats[1].BaseStat
	wildDefence := wildPokemon.Stats[2].BaseStat
	wildAttackDamage := 0
	if wildPokemon.Stats[1].BaseStat-playerPokemon.Stats[2].BaseStat > 0 {
		wildAttackDamage = wildPokemon.Stats[1].BaseStat - playerPokemon.Stats[2].BaseStat
	}

	turn := playerPokemon.Name

	if playerAttackDamage < 0 {
		fmt.Println("Starting battle...")
		time.Sleep(time.Second * 1)
		fmt.Printf("%v was unable to do damage to %v, %v won the battle\n", playerPokemon.Name, wildPokemon.Name, wildPokemon.Name)
		return false
	}
	fmt.Println("")
	fmt.Printf("%v:\n", playerPokemon.Name)
	fmt.Printf("-health: %v\n", playerHealth)
	fmt.Printf("-attack: %v\n", playerAttack)
	fmt.Printf("-defence: %v\n", playerDefence)

	fmt.Println("")
	fmt.Printf("%v:\n", wildPokemon.Name)
	fmt.Printf("-health: %v\n", wildHealth)
	fmt.Printf("-attack: %v\n", wildAttack)
	fmt.Printf("-defence: %v\n", wildDefence)

	for playerHealth > 0 && wildHealth > 0 {

		if turn == playerPokemon.Name {
			time.Sleep(time.Millisecond * 500)
			fmt.Println("")
			fmt.Printf("%v attacking ...\n", playerPokemon.Name)
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("%v did %v damage\n", playerPokemon.Name, playerAttackDamage)
			wildHealth -= playerAttackDamage

			turn = wildPokemon.Name

		} else {
			time.Sleep(time.Millisecond * 500)
			fmt.Println("")
			fmt.Printf("%v attacking ...\n", wildPokemon.Name)
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("%v did %v damage\n", wildPokemon.Name, wildAttackDamage)
			playerHealth -= wildAttackDamage

			turn = playerPokemon.Name
		}

	}

	if playerHealth > 0 {
		return true
	} else {
		return false
	}

}
