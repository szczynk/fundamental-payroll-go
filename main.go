package main

import (
	"fundamental-payroll-go/handler"
	"fundamental-payroll-go/repository"
	"fundamental-payroll-go/usecase"
)

func main() {
	employeeRepo := repository.NewEmployeeRepository()
	employeeUC := usecase.NewEmployeeUsecase(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeUC)

	salaryRepo := repository.NewSalaryRepository()
	salaryUC := usecase.NewSalaryUsecase(salaryRepo)
	salaryHandler := handler.NewSalaryHandler(salaryUC)

	payrollRepo := repository.NewPayrollRepository()
	payrollUC := usecase.NewPayrollUsecase(employeeRepo, payrollRepo, salaryRepo)
	payrollHandler := handler.NewPayrollHandler(payrollUC)

	handler.Menu(employeeHandler, payrollHandler, salaryHandler)
}
