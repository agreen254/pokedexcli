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
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " BULBAsaur CHARIZarddd  Squirtle",
			expected: []string{"bulbasaur", "charizarddd", "squirtle"},
		},
		{
			input:    "      hi       ",
			expected: []string{"hi"},
		},
		{
			input:    "i        E",
			expected: []string{"i", "e"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Fatalf("unequal lengths: expected %d, actual %d", len(c.expected), len(actual))
		}

		for i := range actual {
			got := actual[i]
			want := c.expected[i]

			if got != want {
				t.Errorf("input: %v\n\tgot: %v\n\twant: %v", c.input, got, want)
			}
		}
	}
}
