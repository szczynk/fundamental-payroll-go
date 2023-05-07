package handler

import (
	"fmt"
	"fundamental-payroll-go/helper"
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/helper/input"
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/usecase"
	"strconv"
	"strings"
)

type employeeHandler struct {
	EmployeeUC usecase.EmployeeUsecase
	Input      *input.InputReader
}

func NewEmployeeHandler(employeeUC usecase.EmployeeUsecase, input *input.InputReader) EmployeeHandler {
	h := new(employeeHandler)
	h.EmployeeUC = employeeUC
	h.Input = input

	return h
}

func (handler *employeeHandler) List() {
	_ = helper.ClearTerminal()

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
	_ = helper.ClearTerminal()

	fmt.Println("Add new employee")

	fmt.Print("Name = ")
	name, err := handler.Input.Scan()
	if err != nil || name == "" {
		fmt.Println(apperrors.ErrEmployeeNameNotValid)
		return
	}

	fmt.Print("Gender (laki-laki/perempuan) = ")
	gender, err := handler.Input.Scan()
	trimmedGender := strings.TrimSpace(gender)
	if err != nil || trimmedGender == "" {
		fmt.Println(apperrors.ErrEmployeeGenderNotValid)
		return
	}

	lowerGender := strings.ToLower(trimmedGender)
	switch lowerGender {
	case "l", "laki-laki":
		gender = "laki-laki"
	case "p", "perempuan":
		gender = "perempuan"
	default:
		fmt.Println(apperrors.ErrEmployeeGenderNotValid)
		return
	}

	fmt.Print("Grade = ")
	gradeStr, err := handler.Input.Scan()
	if err != nil {
		fmt.Println(apperrors.ErrEmployeeGradeNotValid)
		return
	}

	grade, err := strconv.ParseInt(strings.TrimSpace(gradeStr), 10, 8)
	if err != nil || grade <= 0 {
		fmt.Println(apperrors.ErrEmployeeGradeNotValid)
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
		Grade:   int8(grade),
		Married: married,
	}

	employee, err := handler.EmployeeUC.Add(employeeRequest)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil add employee with id", employee.ID)
	}
}
