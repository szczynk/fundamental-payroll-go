//go:generate mockery --output=../mocks --name SalaryUsecase
package usecase

import "fundamental-payroll-go/model"

type SalaryUsecase interface {
	List() ([]model.SalaryMatrix, error)
}
