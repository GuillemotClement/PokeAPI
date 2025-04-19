package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	cleanInput := strings.Trim(text, " ")
	lowerInput := strings.ToLower(cleanInput)
	stringText := strings.Split(lowerInput, " ")
	return stringText
}

func main() {
	fmt.Print("Hello, World!")
	cleanInput("text string")
}
