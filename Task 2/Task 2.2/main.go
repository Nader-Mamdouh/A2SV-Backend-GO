package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(input string) bool {
	input = strings.ToLower(input)
	var cleaned strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned.WriteRune(r)
		}
	}

	s := cleaned.String()
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	palindrome := isPalindrome("anna")
	fmt.Println("Is palindrome:", palindrome)
}