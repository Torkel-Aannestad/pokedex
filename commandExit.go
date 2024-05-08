package main

import "os"

func commandExit(*Config) error {
	os.Exit(0)
	return nil
}
