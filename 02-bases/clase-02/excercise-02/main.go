package main

import (
	"fmt"
)

func calculateAverage(grades ...float64) float64 {
	sum := 0.0
	count := 0

	for _, grade := range grades {
		if grade < 0 {
			fmt.Println("Error: Negative grades are not allowed.")
			return 0.0
		}
		sum += grade
		count++
	}

	if count == 0 {
		fmt.Println("Error: No grades to calculate the average.")
		return 0.0
	}

	average := float64(sum) / float64(count)
	return average
}

func main() {

	grades := []float64{4.5, 4.3, 3.2, 2.0, 5.0}

	average := calculateAverage(grades...)

	if average != 0.0 {
		fmt.Printf("The average of the grades is: %.2f\n", average)
	}
}
