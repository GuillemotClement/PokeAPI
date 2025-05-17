package main

import (
	"bufio"
	"fmt"
	"os"
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

func StartRepl() {
	// creation du reader
	reader := bufio.NewScanner(os.Stdin)
	// creation d'une boucle pour recup les input
	for {
		// print pour inviter de commande
		fmt.Print("Pokedex > ")
		// on viens scanner l'input
		reader.Scan()

		// clean input recuperer
		// reader.Text() => recupere l'input scanner
		words := CleanInput(reader.Text())
		// gestion erreur
		if len(words) == 0 {
			// si rien on redemande un nouvel input
			continue
		}

		// on viens recuperer le premier mot de l'input
		commandName := words[0]

		// retour pour user
		fmt.Printf("Your commande was: %s\n", commandName)
	}

}
