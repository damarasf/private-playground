package perusahaan

type VP struct {
	Subordinate []Employee
}

func (vp VP) GetSalary() int {
	return 20
}

func (vp VP) TotalDivisonSalary() int {
	var total int
	for _, employee := range vp.Subordinate {
		totalSalaryDivision := employee.TotalDivisonSalary()
		total += totalSalaryDivision
	}

	totalSalary := vp.GetSalary() + total

	return totalSalary
}
