package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/torkelaannestad/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name     string
	desc     string
	callback func(*Config, ...string) error
}

type Config struct {
	NextUrl         *string
	PreviousUrl     *string
	PokeapiClient   pokeapi.Client
	Pokedex         Pokedex
	CurrentLocation *pokeapi.LocationDetails
}

func startRepl() {
	cfg := &Config{
		PokeapiClient: pokeapi.NewClient(time.Minute, time.Second*15),
		Pokedex:       NewPokedex(),
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex >")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Println("Please provide a command")
			continue
		}
		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		commands := getCommands()
		cmd, exists := commands[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown commmand")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			desc:     "Available Commands",
			callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			desc:     "Exit Pokedex",
			callback: commandExit,
		},
		"map": {
			name:     "map",
			desc:     "List of next locations",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			desc:     "List of previous locations",
			callback: commandMapb,
		},
		"explore": {
			name:     "explore {location name}",
			desc:     "Explores a location by name",
			callback: commandExplore,
		},
		"catch": {
			name:     "catch {pokemon name} {your selected pokemon}",
			desc:     "Throw a pokeball and catch",
			callback: commandCatch,
		},
		"attack": {
			name:     "attack {pokemon name} {your selected pokemon}",
			desc:     "attack and catch a pokemon. Select a pokemon from your pokedex to enter the battle",
			callback: commandAttack,
		},
		"inspect": {
			name:     "inspect {pokemon name}",
			desc:     "See pokemons stats",
			callback: commandInspect,
		},
		"pokedex": {
			name:     "Pokedex",
			desc:     "View your pokemons",
			callback: commandPokedex,
		},
	}
}
