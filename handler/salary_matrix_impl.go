package handler

import (
	"fmt"
	"fundamental-payroll-go/helper"
	"fundamental-payroll-go/helper/input"
	"fundamental-payroll-go/usecase"
)

type salaryHandler struct {
	SalaryUC usecase.SalaryUsecase
	Input    *input.InputReader
}

func NewSalaryHandler(salaryUC usecase.SalaryUsecase, input *input.InputReader) SalaryHandler {
	h := new(salaryHandler)
	h.SalaryUC = salaryUC
	h.Input = input

	return h
}

func (handler *salaryHandler) List() {
	_ = helper.ClearTerminal()

	salaries, err := handler.SalaryUC.List()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("|--------|-------|---------------|---------------|---------------|----------------|\n")
		fmt.Printf("| ID\t | Grade | Basic Salary\t | Pay Cut\t | Allowance\t | Head of Family |\n")
		fmt.Printf("|--------|-------|---------------|---------------|---------------|----------------|\n")

		for _, salary := range salaries {
			if salary.PayCut == 0 {
				fmt.Printf(
					"| %d\t | %d\t | %d\t | %d\t\t | %d\t | %d\t  |\n",
					salary.ID, salary.Grade, salary.BasicSalary, salary.PayCut, salary.Allowance, salary.HoF,
				)
			} else {
				fmt.Printf(
					"| %d\t | %d\t | %d\t | %d\t | %d\t | %d\t  |\n",
					salary.ID, salary.Grade, salary.BasicSalary, salary.PayCut, salary.Allowance, salary.HoF,
				)
			}

		}

		fmt.Printf("|--------|-------|---------------|---------------|---------------|----------------|\n")
	}
}
