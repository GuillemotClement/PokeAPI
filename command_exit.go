package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	// permet de quitter le programme avec un code 0
	os.Exit(0)
	return nil
}
