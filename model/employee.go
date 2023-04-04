package model

type Employee struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Grade   int8   `json:"grade"`
	Married bool   `json:"married"`
}

type EmployeeRequest struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Grade   int8   `json:"grade"`
	Married bool   `json:"married"`
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
