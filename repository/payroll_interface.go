//go:generate mockery --output=../mocks --name PayrollRepository
package repository

import "fundamental-payroll-go/model"

type PayrollRepository interface {
	List() ([]model.Payroll, error)
	Add(payroll *model.Payroll) (*model.Payroll, error)
	Detail(id int64) (*model.Payroll, error)
}
