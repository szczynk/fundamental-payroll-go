package usecase

import (
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/repository"
	"strings"
)

type payrollUsecase struct {
	EmployeeRepo repository.EmployeeRepository
	PayrollRepo  repository.PayrollRepository
	SalaryRepo   repository.SalaryRepository
}

func NewPayrollUsecase(
	employeeRepo repository.EmployeeRepository,
	payrollRepo repository.PayrollRepository,
	salaryRepo repository.SalaryRepository,
) PayrollUsecase {
	return &payrollUsecase{
		EmployeeRepo: employeeRepo,
		PayrollRepo:  payrollRepo,
		SalaryRepo:   salaryRepo,
	}
}

func (uc *payrollUsecase) List() []model.Payroll {
	return uc.PayrollRepo.List()
}

func (uc *payrollUsecase) Add(req *model.PayrollRequest) (*model.Payroll, error) {
	employee, err := uc.EmployeeRepo.Detail(req.EmployeeID)
	if err != nil {
		return &model.Payroll{}, err
	}

	var (
		basicSalary      int64
		payCut           int64
		additionalSalary int64
	)

	for _, salaryMatrix := range uc.SalaryRepo.List() {
		if salaryMatrix.Grade == employee.Grade {
			basicSalary = salaryMatrix.BasicSalary
			payCut = salaryMatrix.PayCut * req.TotalHariTidakMasuk

			additionalSalary = salaryMatrix.Allowance * req.TotalHariMasuk
			if strings.Contains(strings.ToLower(employee.Gender), "laki-laki") && employee.Married {
				additionalSalary += salaryMatrix.HoF
			}
		}
	}

	payroll := &model.Payroll{
		BasicSalary:      basicSalary,
		PayCut:           payCut,
		AdditionalSalary: additionalSalary,
		Employee:         *employee,
	}

	return uc.PayrollRepo.Add(payroll)
}

func (uc *payrollUsecase) Detail(id int64) (*model.Payroll, error) {
	return uc.PayrollRepo.Detail(id)
}
