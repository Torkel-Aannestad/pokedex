package main

import (
	"time"

	"github.com/torkelaannestad/pokedex/internal/pokeapi"
)

func main() {
	config := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 1*time.Minute),
	}

	startRepl(config)
}
