package main

import (
	"fmt"
	"testing"
)

func TestWordFrequencyCount(t *testing.T) {
	testInput := "go go go well , well . . ."
	expected := map[string]int{
		"go":   3,
		"well": 2,
	}

	result := getFrequency(testInput)

	passed := true
	for word, count := range expected {
		if result[word] != count {
			fmt.Printf("Expected %s: %d, got: %d\n", word, count, result[word])
			passed = false
		}
	}
	if passed {
		fmt.Println("Test Passed!")
	}
}
