package handler

import (
	"fmt"
	"fundamental-payroll-go/helper"
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/usecase"
)

type payrollHandler struct {
	PayrollUC usecase.PayrollUsecase
}

func NewPayrollHandler(
	payrollUC usecase.PayrollUsecase,
) PayrollHandler {
	return &payrollHandler{
		PayrollUC: payrollUC,
	}
}

func (handler *payrollHandler) List() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	payrolls, err := handler.PayrollUC.List()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("|--------|---------------|---------------|-------------------|-------------|\n")
		fmt.Printf("| ID\t | Basic Salary\t | Pay Cut\t | Additional Salary | Employee ID |\n")
		fmt.Printf("|--------|---------------|---------------|-------------------|-------------|\n")

		for _, payroll := range payrolls {
			if payroll.PayCut == 0 {
				fmt.Printf(
					"| %d\t | %d\t | %d\t\t | %d\t     | %d\t   |\n",
					payroll.ID, payroll.BasicSalary, payroll.PayCut, payroll.AdditionalSalary, payroll.Employee.ID,
				)
			} else if payroll.AdditionalSalary == 0 {
				fmt.Printf(
					"| %d\t | %d\t | %d\t | %d\t\t     | %d\t   |\n",
					payroll.ID, payroll.BasicSalary, payroll.PayCut, payroll.AdditionalSalary, payroll.Employee.ID,
				)
			} else {
				fmt.Printf(
					"| %d\t | %d\t | %d\t | %d\t     | %d\t   |\n",
					payroll.ID, payroll.BasicSalary, payroll.PayCut, payroll.AdditionalSalary, payroll.Employee.ID,
				)
			}

		}
		fmt.Printf("|--------|---------------|---------------|-------------------|-------------|\n")
	}
}

func (handler *payrollHandler) Add() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Add new payroll")

	fmt.Print("Employee ID = ")
	var employeeID int64
	fmt.Scanln(&employeeID)
	if employeeID <= 0 {
		fmt.Println(apperrors.ErrEmployeeIdNotValid)
		return
	}

	fmt.Print("TotalHariMasuk = ")
	var totalHariMasuk int64
	fmt.Scanln(&totalHariMasuk)
	if totalHariMasuk < 0 {
		fmt.Println(apperrors.ErrPresentDayNotValid)
		return
	}

	fmt.Print("TotalHariTidakMasuk = ")
	var totalHariTidakMasuk int64
	fmt.Scanln(&totalHariTidakMasuk)
	if totalHariTidakMasuk < 0 {
		fmt.Println(apperrors.ErrAbsentDayNotValid)
		return
	}

	payrollRequest := &model.PayrollRequest{
		EmployeeID:          employeeID,
		TotalHariMasuk:      totalHariMasuk,
		TotalHariTidakMasuk: totalHariTidakMasuk,
	}

	payroll, err := handler.PayrollUC.Add(payrollRequest)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil add payroll with id", payroll.ID)
		fmt.Println()
		fmt.Println("|--------------------------------------------------------------------------------|")
		fmt.Println("| Salary Slip\t\t\t\t\t\t\t\t\t |")
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		fmt.Printf(
			"| Payroll ID\t\t | %d\t\t | Nama Employee\t | %s\t |\n",
			payroll.ID, payroll.Employee.Name)
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		fmt.Println("| Penghasilan\t\t | Jumlah\t | Potongan\t\t | Jumlah\t |")
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		if payroll.PayCut == 0 {
			fmt.Printf(
				"| Basic Salary\t\t | %d\t | Pay Cut\t\t | %d\t\t |\n",
				payroll.BasicSalary, payroll.PayCut,
			)
		} else {
			fmt.Printf(
				"| Basic Salary\t\t | %d\t | Pay Cut\t\t | %d\t |\n",
				payroll.BasicSalary, payroll.PayCut,
			)
		}
		fmt.Printf("| Additional Salary\t | %d\t | \t\t\t | \t\t |\n", payroll.AdditionalSalary)
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		if payroll.PayCut == 0 {
			fmt.Printf(
				"| Jumlah Salary\t\t | %d\t | Jumlah Potongan\t | %d\t\t |\n",
				payroll.BasicSalary+payroll.AdditionalSalary, payroll.PayCut,
			)
		} else {
			fmt.Printf("| Jumlah Salary\t\t | %d\t | Jumlah Potongan\t | %d\t |\n",
				payroll.BasicSalary+payroll.AdditionalSalary, payroll.PayCut,
			)
		}
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		fmt.Printf("| Jumlah gaji bersih\t\t\t | %d\t\t\t\t |\n",
			payroll.BasicSalary+payroll.AdditionalSalary-payroll.PayCut,
		)
		fmt.Println("|----------------------------------------|---------------------------------------|")
	}
}

func (handler *payrollHandler) Detail() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Payroll Detail")

	fmt.Print("Payroll ID = ")
	var payrollID int64
	fmt.Scanln(&payrollID)
	if payrollID <= 0 {
		fmt.Println(apperrors.ErrPayrollIdNotValid)
		return
	}

	payroll, err := handler.PayrollUC.Detail(payrollID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil add payroll with id", payroll.ID)
		fmt.Println()
		fmt.Println("|--------------------------------------------------------------------------------|")
		fmt.Println("| Salary Slip\t\t\t\t\t\t\t\t\t |")
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		fmt.Printf(
			"| Payroll ID\t\t | %d\t\t | Nama Employee\t | %s\t |\n",
			payroll.ID, payroll.Employee.Name)
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		fmt.Println("| Penghasilan\t\t | Jumlah\t | Potongan\t\t | Jumlah\t |")
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		if payroll.PayCut == 0 {
			fmt.Printf(
				"| Basic Salary\t\t | %d\t | Pay Cut\t\t | %d\t\t |\n",
				payroll.BasicSalary, payroll.PayCut,
			)
		} else {
			fmt.Printf(
				"| Basic Salary\t\t | %d\t | Pay Cut\t\t | %d\t |\n",
				payroll.BasicSalary, payroll.PayCut,
			)
		}
		fmt.Printf("| Additional Salary\t | %d\t | \t\t\t | \t\t |\n", payroll.AdditionalSalary)
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		if payroll.PayCut == 0 {
			fmt.Printf(
				"| Jumlah Salary\t\t | %d\t | Jumlah Potongan\t | %d\t\t |\n",
				payroll.BasicSalary+payroll.AdditionalSalary, payroll.PayCut,
			)
		} else {
			fmt.Printf("| Jumlah Salary\t\t | %d\t | Jumlah Potongan\t | %d\t |\n",
				payroll.BasicSalary+payroll.AdditionalSalary, payroll.PayCut,
			)
		}
		fmt.Println("|------------------------|---------------|-----------------------|---------------|")
		fmt.Printf("| Jumlah gaji bersih\t\t\t | %d\t\t\t\t |\n",
			payroll.BasicSalary+payroll.AdditionalSalary-payroll.PayCut,
		)
		fmt.Println("|----------------------------------------|---------------------------------------|")
	}
}
