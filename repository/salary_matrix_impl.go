package repository

import (
	"fundamental-payroll-go/model"
)

type salaryRepository struct{}

func NewSalaryRepository() SalaryRepository {
	return new(salaryRepository)
}

func (repo *salaryRepository) List() []model.SalaryMatrix {
	return model.SalaryMatrices
}
