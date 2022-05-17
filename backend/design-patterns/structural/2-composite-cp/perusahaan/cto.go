package perusahaan

import "fmt"

type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

func (c CTO) TotalDivisonSalary() int {
	var total int
	for _, employee := range c.Subordinate {
		totalSalaryDivision := employee.TotalDivisonSalary()
		total += totalSalaryDivision
	}

	totalSalary := c.GetSalary() + total

	fmt.Printf("Total salary CTO: %d\n", totalSalary)

	return totalSalary
}
