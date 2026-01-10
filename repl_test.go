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

// func TestCommandHelp(t *testing.T) {
// 	cases := []struct {
// 		caseName string
// 		input    string
// 		expected string
// 	}{
// 		{
// 			caseName: "normal operation",
// 			input:    "help",
// 			expected: `
// 	Welcome to the Pokedex!
// Usage:

// help: Displays a help message
// exit: Exit the Pokedex
// 	`,
// 		},
// 	}

// 	for _, c := range cases {
// 		err := commandHelp(*config)
// 		if err != nil {
// 			t.Errorf("%v: failing. expected: %v, error: %v", c.caseName, c.expected, err)
// 		}
// 	}
// }

// func TestCommandExit(t *testing.T) {
// 	cases := []struct {
// 		caseName string
// 		input    string
// 		expected string
// 	}{
// 		{
// 			caseName: "normal exit",
// 			input:    "exit",
// 			expected: `Closing the Pokedex... Goodbye!`,
// 		},
// 	}

// 	for _, c := range cases {
// 		err := commandHelp(&config.Config{})
// 		if err != nil {
// 			t.Errorf("%v: failing. expected: %v, error: %v", c.caseName, c.expected, err)
// 		}
// 	}
// }
