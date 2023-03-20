package usecase

import (
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/repository"
)

type salaryUsecase struct {
	SalaryRepo repository.SalaryRepository
}

func NewSalaryUsecase(salaryRepo repository.SalaryRepository) SalaryUsecase {
	return &salaryUsecase{
		SalaryRepo: salaryRepo,
	}
}

func (uc *salaryUsecase) List() []model.SalaryMatrix {
	return uc.SalaryRepo.List()
}
