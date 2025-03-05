package main

import "testing"

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
			input:    "test string",
			expected: []string{"test", "string"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "Test test hello HelLO WORLD",
			expected: []string{"test", "test", "hello", "hello", "world"},
		},
	}
	for _, c := range cases {
		testCase := cleanInput(c.input)
		if len(testCase) != len(c.expected) {
			t.Errorf("Slice length is wrong. Test length is %v, epected length is %v", len(testCase), len(c.expected))
			continue
		}
		for i, word := range testCase {
			if word != c.expected[i] {
				t.Errorf("Slice of strings does not match input string")
			}
		}
	}
}
