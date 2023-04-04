package repository

import (
	"database/sql"
	"fundamental-payroll-go/config/db"
	"fundamental-payroll-go/model"
)

type salaryPgxRepository struct {
	db *sql.DB
}

func NewSalaryPgxRepository(db *sql.DB) SalaryRepository {
	return &salaryPgxRepository{
		db: db,
	}
}

func (repo *salaryPgxRepository) List() ([]model.SalaryMatrix, error) {
	var salaries []model.SalaryMatrix
	var err error

	ctx, cancel := db.NewContext()
	defer cancel()

	sqlQuery := "SELECT id, grade, basic_salary, pay_cut, allowance, head_of_family FROM salaries ORDER BY id ASC"
	rows, err := repo.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return salaries, err
	}
	defer rows.Close()

	for rows.Next() {
		var salary model.SalaryMatrix

		err = rows.Scan(&salary.ID, &salary.Grade, &salary.BasicSalary, &salary.PayCut, &salary.Allowance, &salary.HoF)
		if err != nil {
			return salaries, err
		}

		salaries = append(salaries, salary)
	}

	err = rows.Err()
	if err != nil {
		return salaries, err
	}

	return salaries, nil
}
