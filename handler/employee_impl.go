package handler

import (
	"fmt"
	"fundamental-payroll-go/helper"
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/usecase"
	"strconv"
	"strings"
)

type employeeHandler struct {
	EmployeeUC usecase.EmployeeUsecase
}

func NewEmployeeHandler(employeeUC usecase.EmployeeUsecase) EmployeeHandler {
	return &employeeHandler{
		EmployeeUC: employeeUC,
	}
}

func (handler *employeeHandler) List() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	employees, err := handler.EmployeeUC.List()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("|--------|-----------------------|---------------|-------|---------------|\n")
		fmt.Printf("| ID\t | Nama\t\t\t | Gender\t | Grade | Married\t |\n")
		fmt.Printf("|--------|-----------------------|---------------|-------|---------------|\n")

		for _, employee := range employees {
			if employee.Married {
				fmt.Printf("| %d\t | %s\t | %s\t | %d\t | %t\t\t |\n", employee.ID, employee.Name, employee.Gender, employee.Grade, employee.Married)
			} else {
				fmt.Printf("| %d\t | %s\t | %s\t | %d\t | %t\t |\n", employee.ID, employee.Name, employee.Gender, employee.Grade, employee.Married)
			}
		}

		fmt.Printf("|--------|-----------------------|---------------|-------|---------------|\n")
	}
}

func (handler *employeeHandler) Add() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Add new employee")

	fmt.Print("Name = ")
	var name string
	fmt.Scanln(&name)
	if name == "" {
		fmt.Println(helper.ErrEmployeeNameNotValid)
		return
	}

	fmt.Print("Gender (laki-laki/perempuan) = ")
	var gender string
	fmt.Scanln(&gender)
	gender = strings.ToLower(strings.TrimSpace(gender))
	if gender == "l" {
		gender = "laki-laki"
	} else if gender == "p" {
		gender = "perempuan"
	}
	if gender != "laki-laki" && gender != "perempuan" {
		fmt.Println(helper.ErrEmployeeGenderNotValid)
		return
	}

	fmt.Print("Grade = ")
	var grade int8
	fmt.Scanln(&grade)
	if grade <= 0 {
		fmt.Println(helper.ErrEmployeeGradeNotValid)
		return
	}

	fmt.Print("Married = ")
	var marriedStr string
	fmt.Scanln(&marriedStr)
	married, err := strconv.ParseBool(marriedStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	employeeRequest := &model.EmployeeRequest{
		Name:    name,
		Gender:  gender,
		Grade:   grade,
		Married: married,
	}

	employee, err := handler.EmployeeUC.Add(employeeRequest)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil add employee with id", employee.ID)
	}
}
