package handler

import (
	"fmt"
	"fundamental-payroll-go/helper"
)

func Menu(
	employeeHandler EmployeeHandler,
	payrollHandler PayrollHandler,
	salaryHandler SalaryHandler,
) {
	helper.ClearTerminal()
	helper.ShowMenuList()

	for {
		var menu int
		fmt.Scanln(&menu)

		if menu == 5 {
			helper.ClearTerminal()
			break
		}

		switch menu {
		default: // case 0 atau selain 0
			helper.ClearTerminal()
			helper.ShowMenuList()
		case 1:
			employeeHandler.Add()
		case 2:
			employeeHandler.List()
		case 3:
			payrollHandler.Add()
		case 4:
			salaryHandler.List()
		case 6:
			payrollHandler.List()
		case 7:
			payrollHandler.Detail()
		}
	}
}
