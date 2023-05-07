//go:generate mockery --output=../mocks --name SalaryRepository
package repository

import "fundamental-payroll-go/model"

type SalaryRepository interface {
	List() ([]model.SalaryMatrix, error)
}
