package main

import (
	"fmt"
	"strings"
	"unicode"
)

func getFrequency(input string) map[string]int {
	var cleaned strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			cleaned.WriteRune(r)
		}
	}
	words := strings.Fields(cleaned.String())
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	return freq
}
func main() {
	// Example usage
	freq_of_words := getFrequency("go go go well , well . . .")
	fmt.Println(freq_of_words)

}
