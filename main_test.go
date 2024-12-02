package main

import "testing"

func TestMatchChar(t *testing.T) {
	tests := []struct {
		pattern  string
		char     string
		expected bool
	}{
		{"apple", "apple", true},
		{"pple", "apple", false},
		{".pple", "apple", true},
		{".....", "apple", true},
	}

	for _, test := range tests {
		if output := matchExactWithWildCard(test.pattern, test.char); output != test.expected {
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
		{"^app", "apple", true},
		{"le$", "apple", true},
		{"^a", "apple", true},
		{".$", "apple", true},
		{"apple$", "tasty apple", true},
		{"^apple", "apple pie", true},
		{"^apple$", "apple", true},
		{"^apple$", "tasty apple", false},
		{"^apple$", "apple pie", false},
		{"app$", "apple", false},
		{"^le", "apple", false},
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
