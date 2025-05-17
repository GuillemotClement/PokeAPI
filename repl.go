package main

import (
	"strings"
)

func CleanInput(text string) []string {
	// on lower la chaine fournis
	output := strings.ToLower(text)
	// conversion de la string en slice. Les chaine vide sont supprimer
	words := strings.Fields(output)
	// on retourne le slice
	return words
}
