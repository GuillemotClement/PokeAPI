package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanInput := strings.Trim(text, " ")
	lowerInput := strings.ToLower(cleanInput)
	stringText := strings.Split(lowerInput, " ")
	return stringText
}

func main() {
	// recupere la saisis clavier d'un user
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		inputClean := input[0]

		fmt.Printf("Your command was: %s\n", inputClean)
	}
}
