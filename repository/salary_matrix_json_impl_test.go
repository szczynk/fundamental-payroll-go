package repository

import (
	"encoding/json"
	"fundamental-payroll-go/model"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type JsonSalaryRepoSuite struct {
	suite.Suite
	repo SalaryRepository
}

func mockSalaryJsonFile(salarys *[]model.SalaryMatrix, dir string, pattern string) (tempFileName string, err error) {
	tempFile, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return "", err
	}

	encoder := json.NewEncoder(tempFile)
	err = encoder.Encode(&salarys)
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

func (s *JsonSalaryRepoSuite) SetupSuite() {
	salarys := []model.SalaryMatrix{
		{ID: 1, Grade: 1, BasicSalary: 5000000, PayCut: 100000, Allowance: 150000, HoF: 1000000},
		{ID: 2, Grade: 2, BasicSalary: 9000000, PayCut: 200000, Allowance: 300000, HoF: 2000000},
		{ID: 3, Grade: 3, BasicSalary: 15000000, PayCut: 400000, Allowance: 600000, HoF: 3000000},
	}

	tempFileName, err := mockSalaryJsonFile(&salarys, "", "test_salary_*.json")
	s.Require().NoError(err)

	repo := NewSalaryJsonRepository(tempFileName)

	s.repo = repo
}

func (s *JsonSalaryRepoSuite) TearDownSuite() {
	model.SalaryMatrices = []model.SalaryMatrix{}
	os.Remove(s.repo.(*salaryJsonRepository).jsonFile)
}

func TestJsonSalaryRepoSuite(t *testing.T) {
	suite.Run(t, new(JsonSalaryRepoSuite))
}

func (s *JsonSalaryRepoSuite) Test_salaryJsonRepository_List() {
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

			s.Equal(tt.wantErr, err != nil, "salaryJsonRepository.List() error = %v, wantErr %v", err, tt.wantErr)
			s.Equal(tt.want, got, "salaryJsonRepository.List() = %v, want %v", got, tt.want)
		})
	}
}

func Test_salaryJsonRepository_decodeJSON(t *testing.T) {
	salarys := []model.SalaryMatrix{
		{ID: 1, Grade: 1, BasicSalary: 5000000, PayCut: 100000, Allowance: 150000, HoF: 1000000},
		{ID: 2, Grade: 2, BasicSalary: 9000000, PayCut: 200000, Allowance: 300000, HoF: 2000000},
		{ID: 3, Grade: 3, BasicSalary: 15000000, PayCut: 400000, Allowance: 600000, HoF: 3000000},
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
			pattern: "test_salary_*.json",
			wantErr: false,
		},
		{
			name:    "fail",
			dir:     "tmp/data",
			pattern: "test_salary_*.json",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonFile, err := mockSalaryJsonFile(&salarys, tt.dir, tt.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("mockJsonFile error = %v, wantErr %v", err, tt.wantErr)
			}

			repo := &salaryJsonRepository{
				jsonFile: jsonFile,
			}
			if err := repo.decodeJSON(); (err != nil) != tt.wantErr {
				t.Errorf("salaryJsonRepository.decodeJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
