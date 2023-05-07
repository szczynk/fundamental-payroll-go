package repository

import (
	"fundamental-payroll-go/model"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InMemoryEmployeeRepoSuite struct {
	suite.Suite
	repo EmployeeRepository
}

func (s *InMemoryEmployeeRepoSuite) SetupSuite() {
	model.Employees = []model.Employee{
		{
			ID:      1,
			Name:    "test1",
			Gender:  "laki-laki",
			Grade:   1,
			Married: true,
		},
	}
	repo := NewEmployeeRepository()
	s.repo = repo
}

func (s *InMemoryEmployeeRepoSuite) TearDownSuite() {
	model.Employees = []model.Employee{}
}

func TestInMemoryEmployeeRepoSuite(t *testing.T) {
	suite.Run(t, new(InMemoryEmployeeRepoSuite))
}

func (s *InMemoryEmployeeRepoSuite) Test_employeeRepository_List() {
	tests := []struct {
		name    string
		want    []model.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			want: []model.Employee{
				{
					ID:      1,
					Name:    "test1",
					Gender:  "laki-laki",
					Grade:   1,
					Married: true,
				},
				{
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
			got, err := s.repo.List()

			if s.Equal(tt.wantErr, err != nil, "employeeRepository.List() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "employeeRepository.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *InMemoryEmployeeRepoSuite) Test_employeeRepository_Add() {
	type args struct {
		newEmployee *model.Employee
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				newEmployee: &model.Employee{
					Name:    "test2",
					Gender:  "laki-laki",
					Grade:   1,
					Married: false,
				},
			},
			want: &model.Employee{
				ID:      2,
				Name:    "test2",
				Gender:  "laki-laki",
				Grade:   1,
				Married: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := s.repo.Add(tt.args.newEmployee)

			if s.Equal(tt.wantErr, err != nil, "employeeRepository.Add() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "employeeRepository.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *InMemoryEmployeeRepoSuite) Test_employeeRepository_Detail() {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 2,
			},
			want: &model.Employee{
				ID:      2,
				Name:    "test2",
				Gender:  "laki-laki",
				Grade:   1,
				Married: false,
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

			if s.Equal(tt.wantErr, err != nil, "employeeRepository.Detail() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "employeeRepository.Detail() = %v, want %v", got, tt.want)
			}
		})
	}
}
