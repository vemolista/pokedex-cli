package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "two words with white space",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "leading white space",
			input:    "  hi",
			expected: []string{"hi"},
		},
		{
			name:     "trailing white space",
			input:    "hi  ",
			expected: []string{"hi"},
		},
		{
			name:     "three words with white space",
			input:    "  hi  ho  hey ",
			expected: []string{"hi", "ho", "hey"},
		},
		{
			name:     "capitalized",
			input:    "HI",
			expected: []string{"hi"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Test case %s... For input %q: lengths mismatch. Expected %d, got %d. Expected: %v, Actual: %v",
				c.name, c.input, len(c.expected), len(actual), c.expected, actual)
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
