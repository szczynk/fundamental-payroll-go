package model

type SalaryMatrix struct {
	ID          int64
	Grade       int8
	BasicSalary int64
	PayCut      int64
	Allowance   int64
	HoF         int64
}

var SalaryMatrices = []SalaryMatrix{
	{
		ID:          1,
		Grade:       1,
		BasicSalary: 5_000_000,
		PayCut:      100_000,
		Allowance:   150_000,
		HoF:         1_000_000,
	},
	{
		ID:          2,
		Grade:       2,
		BasicSalary: 9_000_000,
		PayCut:      200_000,
		Allowance:   300_000,
		HoF:         2_000_000,
	},
	{
		ID:          3,
		Grade:       3,
		BasicSalary: 15_000_000,
		PayCut:      400_000,
		Allowance:   600_000,
		HoF:         3_000_000,
	},
}
