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
		// getCommande() permet de retourner le registry de commande
		// [commandName] => on accede a la valeur de cette map avec la cle commandMap
		// en gros, les crochet permet de selectionner la cle que l'on souhaite

		// command => contient la valeur associer a la cle du map OU ou la valeur zero si existe pas ("", 0, false, nil, ..)
		// exists => contient un boolean qui indique si la cle existe reelement dans le map
		command, exists := getCommands()[commandName]
		// si la commande existe bien
		if exists {
			// permet de gerer l'erreur de la fonction callback
			// on execute la commande vu qu'elle existe dans mon map
			err := command.callback()
			// gestion d'erreur retourner par la callback
			if err != nil {
				fmt.Println(err)
			}
			// on retourne faire une iteration => on redemande un nouvel input
			continue
		} else {
			// sinon si la cle existe pas on affiche un message et on redemande une commande
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

// definis le typage pour les commande
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// fonction qui permet de definir le registry de commande
func getCommands() map[string]cliCommand {
	// on retourne le map lors de l'appel de la fonction
	// on type avec le struct => chaque commande doit correspondre au struct
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
