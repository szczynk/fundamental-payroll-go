package model

type Payroll struct {
	ID               int64    `json:"id"`
	BasicSalary      int64    `json:"basic_salary"`
	PayCut           int64    `json:"pay_cut"`
	AdditionalSalary int64    `json:"additional_salary"`
	Employee         Employee `json:"employee"`
	EmployeeID       int64    `json:"employee_id"`
}

type PayrollRequest struct {
	EmployeeID          int64 `json:"employee_id"`
	TotalHariMasuk      int64 `json:"total_hari_masuk"`
	TotalHariTidakMasuk int64 `json:"total_hari_tidak_masuk"`
}

var Payrolls []Payroll

// var Payrolls []Payroll = []Payroll{
// 	{
// 		ID:               1,
// 		BasicSalary:      5_000_000,
// 		PayCut:           300_000,
// 		AdditionalSalary: 2_550_000,
// 		Employee: Employee{
// 			ID:      1,
// 			Name:    "Jonytr",
// 			Gender:  "laki-laki",
// 			Grade:   1,
// 			Married: false,
// 		},
// 	},
// 	{
// 		ID:               2,
// 		BasicSalary:      5_000_000,
// 		PayCut:           2_000_000,
// 		AdditionalSalary: 0,
// 		Employee: Employee{
// 			ID:      2,
// 			Name:    "Jonytr",
// 			Gender:  "laki-laki",
// 			Grade:   1,
// 			Married: false,
// 		},
// 	},
// 	{
// 		ID:               3,
// 		BasicSalary:      5_000_000,
// 		PayCut:           2_000_000,
// 		AdditionalSalary: 1_000_000,
// 		Employee: Employee{
// 			ID:      3,
// 			Name:    "Jonytr",
// 			Gender:  "laki-laki",
// 			Grade:   1,
// 			Married: true,
// 		},
// 	},
// 	{
// 		ID:               4,
// 		BasicSalary:      15_000_000,
// 		PayCut:           0,
// 		AdditionalSalary: 15_000_000,
// 		Employee: Employee{
// 			ID:      4,
// 			Name:    "Jonytr",
// 			Gender:  "laki-laki",
// 			Grade:   3,
// 			Married: true,
// 		},
// 	},
// 	{
// 		ID:               5,
// 		BasicSalary:      15_000_000,
// 		PayCut:           8_000_000,
// 		AdditionalSalary: 0,
// 		Employee: Employee{
// 			ID:      5,
// 			Name:    "Jonytr",
// 			Gender:  "laki-laki",
// 			Grade:   3,
// 			Married: false,
// 		},
// 	},
// }
