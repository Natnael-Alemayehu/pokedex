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
			input:    "  my text",
			expected: []string{"my", "text"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length of the slice of actual: %v and expected: %v doesn't match.", len(actual), len(c.expected))
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Word doesn't match expected word\ngot: %v \nexpected: %v", word, expectedWord)
				t.Fail()
			}
		}
	}
}
