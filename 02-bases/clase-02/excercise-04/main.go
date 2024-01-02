package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation(operationType string) (func(...int) (float64, error), error) {
	switch operationType {
	case minimum:
		return calculateMinimum, nil
	case average:
		return calculateAverage, nil
	case maximum:
		return calculateMaximum, nil
	default:
		return nil, errors.New("Undefined operation type")
	}
}

func calculateMinimum(numbers ...int) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("No numbers provided")
	}

	minValue := float64(numbers[0])
	for _, num := range numbers[1:] {
		if float64(num) < minValue {
			minValue = float64(num)
		}
	}

	return minValue, nil
}

func calculateAverage(numbers ...int) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("No numbers provided")
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	averageValue := float64(sum) / float64(len(numbers))
	return averageValue, nil
}

func calculateMaximum(numbers ...int) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("No numbers provided")
	}

	maxValue := float64(numbers[0])
	for _, num := range numbers[1:] {
		if float64(num) > maxValue {
			maxValue = float64(num)
		}
	}

	return maxValue, nil
}

func main() {
	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	if err != nil {
		fmt.Println(err)
		return
	}

	minValue, err := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue, err := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue, err := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Minimum Value: %.2f\n", minValue)
	fmt.Printf("Average Value: %.2f\n", averageValue)
	fmt.Printf("Maximum Value: %.2f\n", maxValue)
}
