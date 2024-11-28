package taxes

import "fmt"

func calculateSalaryTaxes(salary float64) float64 {
	var taxes float64

	if salary > 50000.00 && salary <= 150000.00 {
		taxes = 0.17 * salary
	} else if salary > 150000.00 {
		taxes = 0.27 * salary
	}
	return taxes
}

func main() {
	var salary	float64
	var	taxes	float64

	fmt.Print("Type the salary to calculate your taxes: ")
	fmt.Scan(&salary)

	taxes = calculateSalaryTaxes(salary)

	fmt.Println("Total salary: US$ ", salary)
	fmt.Println("Total taxes over salary: US$ ", taxes)
}