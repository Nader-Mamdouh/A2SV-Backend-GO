package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateAverage(grades map[string]float64) float64 {
	sum := 0.0
	count := len(grades)
	if count == 0 {
		return 0
	}
	for _, grade := range grades {
		sum += grade
	}
	average := sum / float64(count)
	return math.Round(average*100) / 100 // Rounds to 2 decimal places
}

func main() {
	// Prompt user for input
	fmt.Println("Welcome to the Student Grade Calculator!")
	fmt.Print("Please enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	studentName := scanner.Text()

	fmt.Print("Enter the number of subjects: ")
	scanner.Scan()
	numSubjectsStr := scanner.Text()
	numSubjects, err := strconv.Atoi(numSubjectsStr)
	if err != nil {
		fmt.Println("Error: Invalid number of subjects. Please enter a valid number.")
		return
	}

	grades := make(map[string]float64)

	// Input subjects and grades
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Enter subject %d name: ", i+1)
		scanner.Scan()
		subject := scanner.Text()

		fmt.Printf("Enter grade for %s: ", subject)
		scanner.Scan()
		gradeStr := scanner.Text()
		grade, err := strconv.ParseFloat(gradeStr, 64)
		if err != nil || grade < 0 || grade > 100 {
			fmt.Println("Error: Invalid grade. Please enter a valid numeric grade (0-100).")
			return
		}
		grades[subject] = grade
	}

	// Calculate average grade
	averageGrade := calculateAverage(grades)

	// Display results
	fmt.Printf("\nStudent Name: %s\n", studentName)
	fmt.Println("Subject Grades:")
	for subject, grade := range grades {
		fmt.Printf("- %s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", averageGrade)
}
