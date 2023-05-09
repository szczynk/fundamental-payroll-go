package repository

import (
	"encoding/json"
	"fundamental-payroll-go/model"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type JsonEmployeeRepoSuite struct {
	suite.Suite
	repo EmployeeRepository
}

func mockEmployeeJsonFile(employees *[]model.Employee, dir string, pattern string) (tempFileName string, err error) {
	tempFile, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return "", err
	}

	encoder := json.NewEncoder(tempFile)
	err = encoder.Encode(&employees)
	if err != nil {
		return "", err
	}

	err = tempFile.Close()
	if err != nil {
		return "", err
	}

	tempFileName = tempFile.Name()
	return tempFileName, nil
}

func (s *JsonEmployeeRepoSuite) SetupSuite() {
	employees := []model.Employee{
		{
			ID:      1,
			Name:    "test1",
			Gender:  "laki-laki",
			Grade:   1,
			Married: true,
		},
	}

	tempFileName, err := mockEmployeeJsonFile(&employees, "", "test_employee_*.json")
	s.Require().NoError(err)

	repo := NewEmployeeJsonRepository(tempFileName)

	s.repo = repo
}

func (s *JsonEmployeeRepoSuite) TearDownSuite() {
	model.Employees = []model.Employee{}
	os.Remove(s.repo.(*employeeJsonRepository).jsonFile)
}

func TestJsonEmployeeRepoSuite(t *testing.T) {
	suite.Run(t, new(JsonEmployeeRepoSuite))
}

func (s *JsonEmployeeRepoSuite) Test_employeeJsonRepository_List() {
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

			s.Equal(tt.wantErr, err != nil, "employeeJsonRepository.List() error = %v, wantErr %v", err, tt.wantErr)
			s.Equal(tt.want, got, "employeeJsonRepository.List() = %v, want %v", got, tt.want)
		})
	}
}

func (s *JsonEmployeeRepoSuite) Test_employeeJsonRepository_Add() {
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

			s.Equal(tt.wantErr, err != nil, "employeeJsonRepository.Add() error = %v, wantErr %v", err, tt.wantErr)
			s.Equal(tt.want, got, "employeeJsonRepository.Add() = %v, want %v", got, tt.want)
		})
	}
}

func (s *JsonEmployeeRepoSuite) Test_employeeJsonRepository_Detail() {
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

			s.Equal(tt.wantErr, err != nil, "employeeJsonRepository.Detail() error = %v, wantErr %v", err, tt.wantErr)
			s.Equal(tt.want, got, "employeeJsonRepository.Detail() = %v, want %v", got, tt.want)
		})
	}
}

func Test_employeeJsonRepository_deencodeJSON(t *testing.T) {
	employees := []model.Employee{
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
	}

	tests := []struct {
		name    string
		dir     string
		pattern string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			pattern: "test_employee_*.json",
			wantErr: false,
		},
		{
			name:    "fail",
			dir:     "tmp/data",
			pattern: "test_employee_*.json",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonFile, err := mockEmployeeJsonFile(&employees, tt.dir, tt.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("mockJsonFile error = %v, wantErr %v", err, tt.wantErr)
			}

			repo := &employeeJsonRepository{
				jsonFile: jsonFile,
			}
			if err := repo.encodeJSON(); (err != nil) != tt.wantErr {
				t.Errorf("employeeJsonRepository.encodeJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := repo.decodeJSON(); (err != nil) != tt.wantErr {
				t.Errorf("employeeJsonRepository.decodeJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
