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

func TestMatchRegex(t *testing.T) {
	tests := []struct {
		regex    string
		str      string
		expected bool
	}{
		{"apple", "apple", true},
		{"ap", "apple", true},
		{"le", "apple", true},
		{"a", "apple", true},
		{".", "apple", true},
		{"apwle", "apple", false},
		{"peach", "apple", false},
	}

	for _, tt := range tests {
		t.Run(tt.regex+"_"+tt.str, func(t *testing.T) {
			got := matchRegex(tt.regex, tt.str)
			if got != tt.expected {
				t.Errorf("matchRegexPattern(%q, %q) = %v; want %v", tt.regex, tt.str, got, tt.expected)
			}
		})
	}

}
