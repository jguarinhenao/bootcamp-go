package main

import "fmt"

const (
	categoryC = "C"
	categoryB = "B"
	categoryA = "A"
)

func calculateSalary(minutesWorked int, category string) float64 {
	const hourlyRateC = 1000.0
	const hourlyRateB = 1500.0
	const hourlyRateA = 3000.0

	var monthlySalary float64

	switch category {
	case categoryC:
		monthlySalary = hourlyRateC * float64(minutesWorked) / 60
	case categoryB:
		monthlySalary = (hourlyRateB * float64(minutesWorked) / 60) + (0.20 * hourlyRateB)
	case categoryA:
		monthlySalary = (hourlyRateA * float64(minutesWorked) / 60) + (0.50 * hourlyRateA)
	default:
		fmt.Println("Invalid category.")
		return 0.0
	}

	return monthlySalary
}

func main() {

	minutesWorked := 1200
	category := categoryB

	salary := calculateSalary(minutesWorked, category)
	if salary != 0.0 {
		fmt.Printf("The monthly salary for category %s is: $%.2f\n", category, salary)
	}
}
