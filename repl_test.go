package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    " Do    you believe in    life after love?",
			expected: []string{"do", "you", "believe", "in", "life", "after", "love?"},
		}, {
			input:    " How        about dem    apples    ",
			expected: []string{"how", "about", "dem", "apples"},
		}, {
			input:    " Charmander Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Fail: actual slice length %v does not match expected slice length %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Fail: actual word %s does not match expected word %s", word, expectedWord)
			}
		}
	}
}
