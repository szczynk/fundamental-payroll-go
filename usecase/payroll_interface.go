package usecase

import "fundamental-payroll-go/model"

type PayrollUsecase interface {
	List() ([]model.Payroll, error)
	Add(req *model.PayrollRequest) (*model.Payroll, error)
	Detail(id int64) (*model.Payroll, error)
}
