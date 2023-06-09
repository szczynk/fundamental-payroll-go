package repository

import (
	"database/sql"
	"fundamental-payroll-go/config/db"
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/model"
)

type payrollPgxRepository struct {
	db *sql.DB
}

func NewPayrollPgxRepository(db *sql.DB) PayrollRepository {
	return &payrollPgxRepository{
		db: db,
	}
}

func (repo *payrollPgxRepository) List() ([]model.Payroll, error) {
	var payrolls []model.Payroll
	var err error

	ctx, cancel := db.NewContext()
	defer cancel()

	sqlQuery := "SELECT payrolls.id, payrolls.basic_salary, payrolls.pay_cut, payrolls.additional_salary, payrolls.employee_id, employees.id, employees.name, employees.gender, employees.grade, employees.married FROM payrolls INNER JOIN employees ON payrolls.employee_id = employees.id ORDER BY payrolls.id ASC"
	stmt, err := repo.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return payrolls, err
	}
	defer rows.Close()

	for rows.Next() {
		var payroll model.Payroll
		var employee model.Employee

		err = rows.Scan(
			&payroll.ID, &payroll.BasicSalary, &payroll.PayCut, &payroll.AdditionalSalary, &payroll.EmployeeID,
			&employee.ID, &employee.Name, &employee.Gender, &employee.Grade, &employee.Married,
		)
		if err != nil {
			return payrolls, err
		}

		payroll.Employee = employee

		payrolls = append(payrolls, payroll)
	}

	err = rows.Err()
	if err != nil {
		return payrolls, err
	}

	return payrolls, nil
}

func (repo *payrollPgxRepository) Add(payroll *model.Payroll) (*model.Payroll, error) {
	newPayroll := new(model.Payroll)
	var employee model.Employee
	var err error

	ctx, cancel := db.NewContext()
	defer cancel()

	sqlQuery := "WITH inserted_payroll AS ( INSERT INTO payrolls (basic_salary, pay_cut, additional_salary, employee_id) VALUES ($1, $2, $3, $4) RETURNING id, basic_salary, pay_cut, additional_salary, employee_id ) SELECT id, name, gender, grade, married FROM inserted_payroll INNER JOIN employees ON inserted_payroll.employee_id = employees.id"
	stmt, err := repo.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, payroll.BasicSalary, payroll.PayCut, payroll.AdditionalSalary, payroll.EmployeeID)
	err = row.Scan(
		&payroll.ID, &payroll.BasicSalary, &payroll.PayCut, &payroll.AdditionalSalary, &payroll.EmployeeID,
		&employee.ID, &employee.Name, &employee.Gender, &employee.Grade, &employee.Married,
	)
	if err != nil {
		return nil, err
	}

	return newPayroll, nil
}

func (repo *payrollPgxRepository) Detail(id int64) (*model.Payroll, error) {
	payroll := new(model.Payroll)
	var employee model.Employee
	var err error

	ctx, cancel := db.NewContext()
	defer cancel()

	sqlQuery := "SELECT payrolls.id, payrolls.basic_salary, payrolls.pay_cut, payrolls.additional_salary, payrolls.employee_id, employees.id, employees.name, employees.gender, employees.grade, employees.married FROM payrolls INNER JOIN employees ON payrolls.employee_id = employees.id WHERE payrolls.id = $1 LIMIT 1"
	stmt, err := repo.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(
		&payroll.ID, &payroll.BasicSalary, &payroll.PayCut, &payroll.AdditionalSalary, &payroll.EmployeeID,
		&employee.ID, &employee.Name, &employee.Gender, &employee.Grade, &employee.Married,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.New(apperrors.ErrPayrollNotFound)
		}
		return nil, err
	}

	payroll.Employee = employee

	return payroll, nil
}
