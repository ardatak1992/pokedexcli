package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  pikacHU  World  NicE ",
			expected: []string{"pikachu", "world", "nice"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test

		if len(actual) != len(c.expected) {
			t.Errorf("Word counts don't match\nActual: %d \nExpected: %d\n", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words don't match.\n Actual: '%s'\n Expected:'%s' \n", word, expectedWord)
			}

			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
