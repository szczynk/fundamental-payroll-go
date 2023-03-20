package usecase

import (
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/repository"
)

type employeeUsecase struct {
	EmployeeRepo repository.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo repository.EmployeeRepository) EmployeeUsecase {
	return &employeeUsecase{
		EmployeeRepo: employeeRepo,
	}
}

func (uc *employeeUsecase) List() []model.Employee {
	return uc.EmployeeRepo.List()
}

func (uc *employeeUsecase) Add(req *model.EmployeeRequest) (*model.Employee, error) {
	employee := &model.Employee{
		Name:    req.Name,
		Gender:  req.Gender,
		Grade:   req.Grade,
		Married: req.Married,
	}
	return uc.EmployeeRepo.Add(employee)
}

func (uc *employeeUsecase) Detail(id int64) (*model.Employee, error) {
	return uc.EmployeeRepo.Detail(id)
}
