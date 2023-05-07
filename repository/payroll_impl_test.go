package repository

import (
	"fundamental-payroll-go/model"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InMemoryPayrollRepoSuite struct {
	suite.Suite
	repo PayrollRepository
}

func (s *InMemoryPayrollRepoSuite) SetupSuite() {
	model.Payrolls = []model.Payroll{
		{
			ID:               1,
			BasicSalary:      5000000,
			PayCut:           400000,
			AdditionalSalary: 3400000,
			EmployeeID:       1,
			Employee: model.Employee{
				ID:      1,
				Name:    "test1",
				Gender:  "laki-laki",
				Grade:   1,
				Married: true,
			},
		},
	}
	repo := NewPayrollRepository()
	s.repo = repo
}

func (s *InMemoryPayrollRepoSuite) TearDownSuite() {
	model.Payrolls = []model.Payroll{}
}

func TestInMemoryPayrollRepoSuite(t *testing.T) {
	suite.Run(t, new(InMemoryPayrollRepoSuite))
}

func (s *InMemoryPayrollRepoSuite) Test_payrollRepository_List() {
	tests := []struct {
		name    string
		want    []model.Payroll
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			want: []model.Payroll{
				{
					ID:               1,
					BasicSalary:      5000000,
					PayCut:           400000,
					AdditionalSalary: 3400000,
					EmployeeID:       1,
					Employee: model.Employee{
						ID:      1,
						Name:    "test1",
						Gender:  "laki-laki",
						Grade:   1,
						Married: true,
					},
				},
				{
					ID:               2,
					BasicSalary:      5000000,
					PayCut:           400000,
					AdditionalSalary: 2400000,
					EmployeeID:       1,
					Employee: model.Employee{
						ID:      2,
						Name:    "test2",
						Gender:  "laki-laki",
						Grade:   1,
						Married: false,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := s.repo.List()

			if s.Equal(tt.wantErr, err != nil, "payrollRepository.List() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "payrollRepository.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *InMemoryPayrollRepoSuite) Test_payrollRepository_Add() {
	type args struct {
		newPayroll *model.Payroll
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Payroll
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				newPayroll: &model.Payroll{
					BasicSalary:      5000000,
					PayCut:           400000,
					AdditionalSalary: 2400000,
					EmployeeID:       1,
					Employee: model.Employee{
						ID:      2,
						Name:    "test2",
						Gender:  "laki-laki",
						Grade:   1,
						Married: false,
					},
				},
			},
			want: &model.Payroll{
				ID:               2,
				BasicSalary:      5000000,
				PayCut:           400000,
				AdditionalSalary: 2400000,
				EmployeeID:       1,
				Employee: model.Employee{
					ID:      2,
					Name:    "test2",
					Gender:  "laki-laki",
					Grade:   1,
					Married: false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := s.repo.Add(tt.args.newPayroll)

			if s.Equal(tt.wantErr, err != nil, "payrollRepository.Add() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "payrollRepository.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *InMemoryPayrollRepoSuite) Test_payrollRepository_Detail() {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Payroll
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 2,
			},
			want: &model.Payroll{
				ID:               2,
				BasicSalary:      5000000,
				PayCut:           400000,
				AdditionalSalary: 2400000,
				EmployeeID:       1,
				Employee: model.Employee{
					ID:      2,
					Name:    "test2",
					Gender:  "laki-laki",
					Grade:   1,
					Married: false,
				},
			},
			wantErr: false,
		},
		{
			name: "failed",
			args: args{
				id: 5,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := s.repo.Detail(tt.args.id)

			if s.Equal(tt.wantErr, err != nil, "payrollRepository.Detail() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "payrollRepository.Detail() = %v, want %v", got, tt.want)
			}
		})
	}
}
