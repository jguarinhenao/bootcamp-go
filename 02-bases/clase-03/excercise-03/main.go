package main

import (
	"fmt"
	"time"
)

// Student structure
type Student struct {
	Name     string
	LastName string
	ID       string
	Date     time.Time
}

// Display prints the details of the student
func (s *Student) Display() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Last Name:", s.LastName)
	fmt.Println("ID:", s.ID)
	fmt.Println("Date:", s.Date.Format("02/01/2006")) // Date format: MM/DD/YYYY
}

func main() {
	// Instantiate a student
	student := Student{
		Name:     "John",
		LastName: "Doe",
		ID:       "123456789",
		Date:     time.Now(),
	}

	// Display the details of the student
	student.Display()
}
