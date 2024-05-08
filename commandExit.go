package main

import (
	"fmt"
	"os"
)

func commandExit(config *config) error {
	fmt.Println()
	fmt.Println("Exiting Pokedex...")
	fmt.Println()
	os.Exit(0)
	return nil
}
