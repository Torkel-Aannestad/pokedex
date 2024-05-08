package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Println()
	fmt.Println("Welcome")
	fmt.Println("Use one of the following commands")
	for _, cmd := range getCommands() {
		fmt.Printf("Name: %v, Description %v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
