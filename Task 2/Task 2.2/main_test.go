package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abba", true},
		{"A man, a plan, a canal, Panama", true},
		{"Hello World", false},
		{"Was it a car or a cat I saw", true},
		{"No 'x' in Nixon", true},
		{"hello", false},
		{"", true},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
		}
	}
}
