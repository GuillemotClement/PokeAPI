package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	// declaration d'un slice de cas de test.
	// chaque cas est une struct avec deux champs : input et exepected
	cases := []struct {
		// declaration du struct
		input    string   // chaine a traiter
		expected []string // resultat attendu
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}
	// on itere le slice de cases
	// c est le cas de l'iteration en cours
	// pour chaque cas de test on appelle la fonction a tester
	for _, c := range cases {
		// on recupere l'input du slice
		actual := CleanInput(c.input)

		// on verifie la taille du resulat
		// si le nombre d'element est different, alors erreur
		// avec le %q on affiche des guillemts pour ameliorer la visibilite
		if len(actual) != len(c.expected) {
			// affiche un message de log
			t.Errorf("Length don't match: %d vs %d", actual, c.expected)
			// permet de continuer les tests
			continue
		}
		// on verifie chaque mot est le meme
		// i correspond au mot
		for i := range actual {
			// on recupere le mot du input
			word := actual[i]
			// on recupere le mot equivalent dans le expected
			expectedWord := c.expected[i]

			// si le mot est different on echoue le test
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
