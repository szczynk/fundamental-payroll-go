package repository

import (
	"database/sql"
	"fundamental-payroll-go/config/db"
	"fundamental-payroll-go/model"
)

type employeePgxRepository struct {
	db *sql.DB
}

func NewEmployeePgxRepository(db *sql.DB) EmployeeRepository {
	return &employeePgxRepository{
		db: db,
	}
}

func (repo *employeePgxRepository) List() ([]model.Employee, error) {
	var employees []model.Employee
	var err error

	ctx, cancel := db.NewContext()
	defer cancel()

	sqlQuery := "SELECT id, name, gender, grade, married FROM employees ORDER BY id ASC"
	rows, err := repo.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return employees, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee model.Employee

		err = rows.Scan(&employee.ID, &employee.Name, &employee.Gender, &employee.Grade, &employee.Married)
		if err != nil {
			return employees, err
		}

		employees = append(employees, employee)
	}

	err = rows.Err()
	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (repo *employeePgxRepository) Add(employee *model.Employee) (*model.Employee, error) {
	newEmployee := new(model.Employee)
	var err error

	ctx, cancel := db.NewContext()
	defer cancel()

	sqlQuery := `
	INSERT INTO employees (name, gender, grade, married) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, name, gender, grade, married
	`
	row := repo.db.QueryRowContext(ctx, sqlQuery, employee.Name, employee.Gender, employee.Grade, employee.Married)
	err = row.Scan(&newEmployee.ID, &newEmployee.Name, &newEmployee.Gender, &newEmployee.Grade, &newEmployee.Married)
	if err != nil {
		return nil, err
	}

	return newEmployee, nil
}

func (repo *employeePgxRepository) Detail(id int64) (*model.Employee, error) {
	employee := new(model.Employee)
	var err error

	ctx, cancel := db.NewContext()
	defer cancel()

	sqlQuery := "SELECT id, name, gender, grade, married FROM employees WHERE id = $1 LIMIT 1"
	row := repo.db.QueryRowContext(ctx, sqlQuery, id)
	err = row.Scan(&employee.ID, &employee.Name, &employee.Gender, &employee.Grade, &employee.Married)
	if err != nil {
		return nil, err
	}

	return employee, nil
}
