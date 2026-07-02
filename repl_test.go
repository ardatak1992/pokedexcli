package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "this is a test    ",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    " TEst the cleanInput Bulbasaur PIKACHU",
			expected: []string{"test", "the", "cleaninput", "bulbasaur", "pikachu"},
		},
		{
			input:    "Charmander        Bulbasaur      PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLength := len(actual)
		expectedLength := len(c.expected)
		if actualLength != expectedLength {
			// Check length of string
			t.Errorf("Mismath in slice lengths:\nActual: %d\nExpected: %d\n", actualLength, expectedLength)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Mismatch in words:\nActual: %s\nExpected: %s\n", word, expectedWord)
			}

		}

	}

}
