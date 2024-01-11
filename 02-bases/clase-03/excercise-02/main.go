package main

import (
	"fmt"
	"time"
)

// Person estructura
type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

// Employee estructura con composición de Person
type Employee struct {
	ID       int
	Position string
	Person   Person
}

// PrintEmployee imprime los campos de un empleado
func (e *Employee) PrintEmployee() {
	fmt.Printf("ID: %d, Position: %s\n", e.ID, e.Position)
	fmt.Printf("Person - ID: %d, Name: %s, Date of Birth: %s\n",
		e.Person.ID, e.Person.Name, e.Person.DateOfBirth.Format("2006-01-02"))
}

func main() {
	// Instanciar una Person
	person := Person{ID: 1, Name: "John Doe", DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)}

	// Instanciar un Employee con composición de Person
	employee := Employee{ID: 101, Position: "Software Engineer", Person: person}

	// Ejecutar el método PrintEmployee() para imprimir los campos del empleado
	employee.PrintEmployee()
}
