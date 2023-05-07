package repository

import (
	"fundamental-payroll-go/model"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InMemorySalaryRepoSuite struct {
	suite.Suite
	repo SalaryRepository
}

func (s *InMemorySalaryRepoSuite) SetupSuite() {
	model.SalaryMatrices = []model.SalaryMatrix{
		{ID: 1, Grade: 1, BasicSalary: 5000000, PayCut: 100000, Allowance: 150000, HoF: 1000000},
		{ID: 2, Grade: 2, BasicSalary: 9000000, PayCut: 200000, Allowance: 300000, HoF: 2000000},
		{ID: 3, Grade: 3, BasicSalary: 15000000, PayCut: 400000, Allowance: 600000, HoF: 3000000},
	}

	repo := NewSalaryRepository()
	s.repo = repo
}

func (s *InMemorySalaryRepoSuite) TearDownSuite() {
	model.SalaryMatrices = []model.SalaryMatrix{}
}

func TestInMemorySalaryRepoSuite(t *testing.T) {
	suite.Run(t, new(InMemorySalaryRepoSuite))
}

func (s *InMemorySalaryRepoSuite) Test_salaryRepository_List() {
	tests := []struct {
		name    string
		want    []model.SalaryMatrix
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			want: []model.SalaryMatrix{
				{ID: 1, Grade: 1, BasicSalary: 5000000, PayCut: 100000, Allowance: 150000, HoF: 1000000},
				{ID: 2, Grade: 2, BasicSalary: 9000000, PayCut: 200000, Allowance: 300000, HoF: 2000000},
				{ID: 3, Grade: 3, BasicSalary: 15000000, PayCut: 400000, Allowance: 600000, HoF: 3000000},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := s.repo.List()

			if s.Equal(tt.wantErr, err != nil, "salaryRepository.List() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "salaryRepository.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
