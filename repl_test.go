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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  hi",
			expected: []string{"hi"},
		},
		{
			input:    "hi  ",
			expected: []string{"hi"},
		},
		{
			input:    "  hi  ho  hey ",
			expected: []string{"hi", "ho", "hey"},
		},
		{
			input:    "HI",
			expected: []string{"hi"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("For input %q: lengths mismatch. Expected %d, got %d. Expected: %v, Actual: %v",
				c.input, len(c.expected), len(actual), c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words not matching expected words")
			}
		}
	}
}
