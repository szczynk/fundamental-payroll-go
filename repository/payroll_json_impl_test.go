package repository

import (
	"encoding/json"
	"fundamental-payroll-go/model"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type JsonPayrollRepoSuite struct {
	suite.Suite
	repo PayrollRepository
}

func mockPayrollJsonFile(payrolls *[]model.Payroll, dir string, pattern string) (tempFileName string, err error) {
	tempFile, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return "", err
	}

	encoder := json.NewEncoder(tempFile)
	err = encoder.Encode(&payrolls)
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

func (s *JsonPayrollRepoSuite) SetupSuite() {
	payrolls := []model.Payroll{
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

	tempFileName, err := mockPayrollJsonFile(&payrolls, "", "test_payroll_*.json")
	s.Require().NoError(err)

	repo := NewPayrollJsonRepository(tempFileName)

	s.repo = repo
}

func (s *JsonPayrollRepoSuite) TearDownSuite() {
	model.Payrolls = []model.Payroll{}
	os.Remove(s.repo.(*payrollJsonRepository).jsonFile)
}

func TestJsonPayrollRepoSuite(t *testing.T) {
	suite.Run(t, new(JsonPayrollRepoSuite))
}

func (s *JsonPayrollRepoSuite) Test_payrollJsonRepository_List() {
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

			s.Equal(tt.wantErr, err != nil, "payrollJsonRepository.List() error = %v, wantErr %v", err, tt.wantErr)
			s.Equal(tt.want, got, "payrollJsonRepository.List() = %v, want %v", got, tt.want)
		})
	}
}

func (s *JsonPayrollRepoSuite) Test_payrollJsonRepository_Add() {
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

			s.Equal(tt.wantErr, err != nil, "payrollJsonRepository.Add() error = %v, wantErr %v", err, tt.wantErr)
			s.Equal(tt.want, got, "payrollJsonRepository.Add() = %v, want %v", got, tt.want)
		})
	}
}

func (s *JsonPayrollRepoSuite) Test_payrollJsonRepository_Detail() {
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

			s.Equal(tt.wantErr, err != nil, "payrollJsonRepository.Detail() error = %v, wantErr %v", err, tt.wantErr)
			s.Equal(tt.want, got, "payrollJsonRepository.Detail() = %v, want %v", got, tt.want)
		})
	}
}

func Test_payrollJsonRepository_deencodeJSON(t *testing.T) {
	payrolls := []model.Payroll{
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
			pattern: "test_payroll_*.json",
			wantErr: false,
		},
		{
			name:    "fail",
			dir:     "tmp/data",
			pattern: "test_payroll_*.json",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonFile, err := mockPayrollJsonFile(&payrolls, tt.dir, tt.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("mockJsonFile error = %v, wantErr %v", err, tt.wantErr)
			}

			repo := &payrollJsonRepository{
				jsonFile: jsonFile,
			}
			if err := repo.encodeJSON(); (err != nil) != tt.wantErr {
				t.Errorf("payrollJsonRepository.encodeJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := repo.decodeJSON(); (err != nil) != tt.wantErr {
				t.Errorf("payrollJsonRepository.decodeJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
