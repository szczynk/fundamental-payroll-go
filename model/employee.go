package model

type Employee struct {
	ID      int64
	Name    string
	Gender  string
	Grade   int8
	Married bool
}

type EmployeeRequest struct {
	Name    string
	Gender  string
	Grade   int8
	Married bool
}

var Employees []Employee

// var Employees []Employee = []Employee{
// 	{
// 		ID:      2,
// 		Name:    "Muhammad Wong",
// 		Gender:  "laki-laki",
// 		Grade:   2,
// 		Married: true,
// 	},
// 	{
// 		ID:      3,
// 		Name:    "Bagus Nuryasin",
// 		Gender:  "laki-laki",
// 		Grade:   3,
// 		Married: true,
// 	},
// }
