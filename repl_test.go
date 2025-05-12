package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   leading and trailing   ",
			expected: []string{"leading", "and", "trailing"},
		},
		{
			input:    "multiple   spaces   between words",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input:    "\t tabs\tand\nnewlines\n",
			expected: []string{"tabs", "and", "newlines"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len(actual) = %d, len(expected) = %d", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual[%d] = %s, expected[%d] = %s", i, word, i, expectedWord)
			}
		}
	}
}
