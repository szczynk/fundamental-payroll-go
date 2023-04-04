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

func (uc *payrollUsecase) List() ([]model.Payroll, error) {
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

	salaries, err := uc.SalaryRepo.List()
	if err != nil {
		return &model.Payroll{}, err
	}

	for _, salary := range salaries {
		if salary.Grade == employee.Grade {
			basicSalary = salary.BasicSalary
			payCut = salary.PayCut * req.TotalHariTidakMasuk

			additionalSalary = salary.Allowance * req.TotalHariMasuk
			if strings.Contains(strings.ToLower(employee.Gender), "laki-laki") && employee.Married {
				additionalSalary += salary.HoF
			}
		}
	}

	payroll := &model.Payroll{
		BasicSalary:      basicSalary,
		PayCut:           payCut,
		AdditionalSalary: additionalSalary,
		Employee:         *employee,
		EmployeeID:       req.EmployeeID,
	}

	return uc.PayrollRepo.Add(payroll)
}

func (uc *payrollUsecase) Detail(id int64) (*model.Payroll, error) {
	return uc.PayrollRepo.Detail(id)
}
