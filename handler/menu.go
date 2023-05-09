package handler

import (
	"errors"
	"fmt"
	"fundamental-payroll-go/helper/input"
	"strconv"
	"strings"
)

type Menu struct {
	employeeH    EmployeeHandler
	payrollH     PayrollHandler
	salaryH      SalaryHandler
	i            *input.InputReader
	clear        func() error
	showMenuList func()
}

func NewMenu(employeeHandler EmployeeHandler, payrollHandler PayrollHandler, salaryHandler SalaryHandler, input *input.InputReader, clear func() error, showMenuList func()) *Menu {
	menu := new(Menu)
	menu.employeeH = employeeHandler
	menu.payrollH = payrollHandler
	menu.salaryH = salaryHandler
	menu.i = input
	menu.clear = clear
	menu.showMenuList = showMenuList

	return menu
}

func (m *Menu) ShowMenu() error {
	err := m.clear()
	if err != nil {
		return err
	}

	m.showMenuList()

	for {
		menuStr, _ := m.i.Scan()
		menu64, err := strconv.ParseInt(strings.TrimSpace(menuStr), 10, 64)
		if err != nil && !errors.Is(err, strconv.ErrSyntax) {
			fmt.Println(err)
			break
		}
		menu := int32(menu64)

		if menu == 5 {
			_ = m.clear()
			break
		}

		switch menu {
		default: // case 0 atau selain 0
			_ = m.clear
			m.showMenuList()
		case 1:
			fmt.Println("Add a new employee")
			m.employeeH.Add()
		case 2:
			fmt.Println("Employee list")
			m.employeeH.List()
		case 3:
			fmt.Println("Add a new payroll")
			m.payrollH.Add()
		case 4:
			fmt.Println("Salary Matrix")
			m.salaryH.List()
		case 6:
			fmt.Println("Payroll list")
			m.payrollH.List()
		case 7:
			fmt.Println("Payroll detail")
			m.payrollH.Detail()
		}
	}

	return nil
}
