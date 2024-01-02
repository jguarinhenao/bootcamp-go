package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func animal(animalType string) (func(float64) float64, error) {
	switch animalType {
	case dog:
		return calculateDogFood, nil
	case cat:
		return calculateCatFood, nil
	case hamster:
		return calculateHamsterFood, nil
	case tarantula:
		return calculateTarantulaFood, nil
	default:
		return nil, errors.New("Undefined animal type")
	}
}

func calculateDogFood(amount float64) float64 {
	return amount * 10
}

func calculateCatFood(amount float64) float64 {
	return amount * 5
}

func calculateHamsterFood(amount float64) float64 {
	return amount * 0.25
}

func calculateTarantulaFood(amount float64) float64 {
	return amount * 0.15
}

func main() {
	animalDog, msg := animal(dog)
	animalCat, msg := animal(cat)

	if msg != nil {
		fmt.Println(msg)
		return
	}

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Printf("Total amount of food needed: %.2f kg\n", amount)
}
