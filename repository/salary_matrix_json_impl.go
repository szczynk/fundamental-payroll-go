package repository

import (
	"encoding/json"
	"fundamental-payroll-go/model"
	"os"
)

type salaryJsonRepository struct {
	jsonFile string
}

func NewSalaryJsonRepository(jsonFilePath string) SalaryRepository {
	r := new(salaryJsonRepository)
	r.jsonFile = jsonFilePath

	return r
}
func (repo *salaryJsonRepository) decodeJSON() error {
	reader, err := os.Open(repo.jsonFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&model.SalaryMatrices)
	if err != nil {
		return err
	}
	return nil
}

func (repo *salaryJsonRepository) List() ([]model.SalaryMatrix, error) {
	err := repo.decodeJSON()
	if err != nil {
		return []model.SalaryMatrix{}, err
	}

	return model.SalaryMatrices, nil
}
