package handler

import (
	"errors"
	"fmt"
	"fundamental-payroll-go/helper"
	"fundamental-payroll-go/helper/input"
	"strconv"
	"strings"
)

func Menu(
	employeeHandler EmployeeHandler,
	payrollHandler PayrollHandler,
	salaryHandler SalaryHandler,
	input *input.InputReader,
) {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	helper.ShowMenuList()

	for {
		menuStr, _ := input.Scan()
		menu64, err := strconv.ParseInt(strings.TrimSpace(menuStr), 10, 32)
		if err != nil && !errors.Is(err, strconv.ErrSyntax) {
			fmt.Println(err)
			break
		}
		menu := int32(menu64)

		if menu == 5 {
			_ = helper.ClearTerminal()
			break
		}

		switch menu {
		default: // case 0 atau selain 0
			_ = helper.ClearTerminal()
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
