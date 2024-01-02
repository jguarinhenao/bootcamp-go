package excerciseone

import (
	"testing"
)

func TestCalculateSalaryTax(t *testing.T) {
	// Caso 1
	salaryCase1 := 45000.0
	expectedFinalSalary1 := 450000.0
	expectedTax1 := 10.0
	finalSalaryCase1, actualTaxCase1 := CalculateSalaryTax(salaryCase1)
	if finalSalaryCase1 != expectedFinalSalary1 || actualTaxCase1 != expectedTax1 {
		t.Errorf("Caso 1 fallido. Salario: %.2f, Salario final esperado: %.2f, Impuesto esperado: %.2f, Salario final actual: %.2f, Impuesto actual: %.2f",
			salaryCase1, expectedFinalSalary1, expectedTax1, finalSalaryCase1, actualTaxCase1)
	}

	// Caso 2
	salaryCase2 := 55000.0
	expectedTax2 := 17.0
	finalSalaryCase2, actualTaxCase2 := CalculateSalaryTax(salaryCase2)
	if actualTaxCase2 != expectedTax2 {
		t.Errorf("Caso 2 fallido. Salario: %.2f, Impuesto esperado: %.2f, Impuesto actual: %.2f",
			salaryCase2, expectedTax2, actualTaxCase2)
	}

	// Caso 3
	salaryCase3 := 160000.0
	expectedTax3 := 27.0
	finalSalaryCase3, actualTaxCase3 := CalculateSalaryTax(salaryCase3)
	if actualTaxCase3 != expectedTax3 {
		t.Errorf("Caso 3 fallido. Salario: %.2f, Impuesto esperado: %.2f, Impuesto actual: %.2f",
			salaryCase3, expectedTax3, actualTaxCase3)
	}
}
