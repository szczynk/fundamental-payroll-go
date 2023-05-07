package repository

import (
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/model"
)

type payrollRepository struct{}

func NewPayrollRepository() PayrollRepository {
	return new(payrollRepository)
}

func (repo *payrollRepository) List() ([]model.Payroll, error) {
	return model.Payrolls, nil
}

func (repo *payrollRepository) getLastID() int64 {
	payrolls, _ := repo.List()

	var tempID int64
	for _, payroll := range payrolls {
		if tempID < payroll.ID {
			tempID = payroll.ID
		}
	}
	return tempID
}

func (repo *payrollRepository) Add(payroll *model.Payroll) (*model.Payroll, error) {
	id := repo.getLastID()

	newPayroll := payroll
	newPayroll.ID = id + 1

	model.Payrolls = append(model.Payrolls, *newPayroll)

	return newPayroll, nil
}

func (repo *payrollRepository) getIndexByID(id int64) (int, error) {
	payrolls, _ := repo.List()

	for i, payroll := range payrolls {
		if id == payroll.ID {
			return i, nil
		}
	}

	return -1, apperrors.New(apperrors.ErrIdNotFound)
}

func (repo *payrollRepository) Detail(id int64) (*model.Payroll, error) {
	payrolls, _ := repo.List()

	index, err := repo.getIndexByID(id)
	if err != nil {
		return nil, err
	}

	payroll := payrolls[index]

	return &payroll, nil
}
