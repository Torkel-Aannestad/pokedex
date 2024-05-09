package main

import "fmt"

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("")
	fmt.Println("Here are the available commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("\t-%v, Description: %v\n", cmd.name, cmd.desc)
	}
	return nil
}
