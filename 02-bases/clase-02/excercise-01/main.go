package excerciseone

import "fmt"

func main() {

	var salary float64 = 51000

	finallySalary, totalTaxes := CalculateSalaryTax(salary)

	fmt.Printf("Final salary: %.2f, with a discount of %.2f%%\n", finallySalary, totalTaxes)

}

func CalculateSalaryTax(salary float64) (finallySalary float64, totalTaxes float32) {
	switch {
	case salary > 150000:
		totalTaxes = 27
	case salary > 50000:
		totalTaxes = 17
	default:
		totalTaxes = 0
	}

	finallySalary = salary - (salary * float64(totalTaxes) / 100)

	return
}
