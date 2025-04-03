package main

import (
	"math"
	"testing"
)

func almostEqual(a, b float64, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		grades   map[string]float64
		expected float64
	}{
		{"Single subject", map[string]float64{"Math": 80}, 80},
		{"Multiple subjects", map[string]float64{"Math": 70, "English": 80, "Science": 90}, 80},
		{"All zeroes", map[string]float64{"Math": 0, "English": 0, "Science": 0}, 0},
		{"Mixed values", map[string]float64{"Physics": 50, "Chemistry": 60, "Biology": 70}, 60},
		{"Decimal values", map[string]float64{"CS": 88.5, "AI": 91.2, "ML": 75.6}, 85.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateAverage(tt.grades)
			if !almostEqual(result, tt.expected, 0.01) { // Allow small differences
				t.Errorf("calculateAverage(%v) = %v; want %v", tt.grades, result, tt.expected)
			}
		})
	}
}
