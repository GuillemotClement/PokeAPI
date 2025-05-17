package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

		// on recupere la commande depuis input de user
		commandName := words[0]

		// on viens checked si la commande existe dans le registry des commandes
		command, exists := getCommands()[commandName]
		// si la commande existe bien
		if exists {
			// permet de gerer l'erreur de la fonction callback
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknow command")
			continue
		}
	}
}

func CleanInput(text string) []string {
	// on lower la chaine fournis
	output := strings.ToLower(text)
	// conversion de la string en slice. Les chaine vide sont supprimer
	words := strings.Fields(output)
	// on retourne le slice
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
