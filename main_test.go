package main

import "testing"

func TestMatchChar(t *testing.T) {
	tests := []struct {
		pattern  byte
		char     byte
		expected bool
	}{
		{'a', 'a', true},
		{'a', 'b', false},
		{'.', 'a', true},
		{'.', 'b', true},
	}

	for _, test := range tests {
		if output := matchChar(test.pattern, test.char); output != test.expected {
			t.Errorf("Test failed: %v inputted, %v expected, %v received", test, test.expected, output)
		}
	}
}
