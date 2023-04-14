package usecase

import (
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/repository"
)

type employeeUsecase struct {
	EmployeeRepo repository.EmployeeRepository
	SalaryRepo   repository.SalaryRepository
}

func NewEmployeeUsecase(
	employeeRepo repository.EmployeeRepository,
	salaryRepo repository.SalaryRepository,
) EmployeeUsecase {
	return &employeeUsecase{
		EmployeeRepo: employeeRepo,
		SalaryRepo:   salaryRepo,
	}
}

func (uc *employeeUsecase) List() ([]model.Employee, error) {
	return uc.EmployeeRepo.List()
}

func (uc *employeeUsecase) Add(req *model.EmployeeRequest) (*model.Employee, error) {
	salaries, err := uc.SalaryRepo.List()
	if err != nil {
		return nil, err
	}

	var isValidGrade bool
	for _, salary := range salaries {
		if req.Grade == salary.Grade {
			isValidGrade = true
			break
		}
	}
	if !isValidGrade {
		return nil, apperrors.New(apperrors.ErrEmployeeGradeNotValid)
	}

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
