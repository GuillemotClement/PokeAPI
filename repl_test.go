package main

import "testing"

func TestCleanInput(t *testing.T) {
	// creation d'un slice de test
	cases := []struct {
		input    string
		expected []string
	}{
		{
			// input de la fonction
			input: " Hello World ",
			// resultat en sortie
			expected: []string{"hello", "world"},
		},
	}
	// ajout d'autre cas a tester

	// iteration sur le contenu
	for _, c := range cases {
		// on passe sur les case definis, et on les passe dans la fonction
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch for input %v: got %v, want %v", c.input, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word mismatch at position %v for input %v: got %v, want %v", i, c.input, word, expectedWord)
			}
		}
	}

}
