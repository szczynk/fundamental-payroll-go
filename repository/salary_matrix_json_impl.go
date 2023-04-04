package repository

import (
	"encoding/json"
	"fundamental-payroll-go/model"
	"os"
)

const salaryJsonFile = "data/salary.json"

type salaryJsonRepository struct{}

func NewSalaryJsonRepository() SalaryRepository {
	return new(salaryJsonRepository)
}

func (repo *salaryJsonRepository) decodeJSON(path string, salaries *[]model.SalaryMatrix) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&salaries)
	if err != nil {
		return err
	}
	return nil
}

func (repo *salaryJsonRepository) List() ([]model.SalaryMatrix, error) {
	err := repo.decodeJSON(salaryJsonFile, &model.SalaryMatrices)
	if err != nil {
		return []model.SalaryMatrix{}, err
	}

	return model.SalaryMatrices, nil
}
