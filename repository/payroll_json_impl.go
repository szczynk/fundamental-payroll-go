package repository

import (
	"encoding/json"
	"errors"
	"fundamental-payroll-go/model"
	"os"
)

const payrollJsonFile = "data/payroll.json"

type payrollJsonRepository struct{}

func NewPayrollJsonRepository() PayrollRepository {
	return new(payrollJsonRepository)
}

func (repo *payrollJsonRepository) encodeJSON(path string, payrolls *[]model.Payroll) error {
	writer, err := os.Create(path)
	if err != nil {
		return err
	}
	defer writer.Close()

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&payrolls)
	if err != nil {
		return err
	}
	return nil
}

func (repo *payrollJsonRepository) decodeJSON(path string, payrolls *[]model.Payroll) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&payrolls)
	if err != nil {
		return err
	}
	return nil
}

func (repo *payrollJsonRepository) List() ([]model.Payroll, error) {
	err := repo.decodeJSON(payrollJsonFile, &model.Payrolls)
	if err != nil {
		return []model.Payroll{}, err
	}

	return model.Payrolls, nil
}

func (repo *payrollJsonRepository) getLastID() (int64, error) {
	payrolls, err := repo.List()

	var tempID int64
	for _, payroll := range payrolls {
		if tempID < payroll.ID {
			tempID = payroll.ID
		}
	}
	return tempID, err
}

func (repo *payrollJsonRepository) Add(payroll *model.Payroll) (*model.Payroll, error) {
	id, err := repo.getLastID()
	if err != nil {
		return &model.Payroll{}, nil
	}

	newPayroll := payroll
	newPayroll.ID = id + 1

	model.Payrolls = append(model.Payrolls, *newPayroll)

	err = repo.encodeJSON(payrollJsonFile, &model.Payrolls)
	if err != nil {
		return &model.Payroll{}, err
	}

	return newPayroll, nil
}

func (repo *payrollJsonRepository) getIndexByID(id int64) (int, error) {
	payrolls, err := repo.List()
	if err != nil {
		return -1, err
	}

	for i, payroll := range payrolls {
		if id == payroll.ID {
			return i, nil
		}
	}

	return -1, errors.New("ID tidak ditemukan")
}

func (repo *payrollJsonRepository) Detail(id int64) (*model.Payroll, error) {
	payrolls, err := repo.List()
	if err != nil {
		return &model.Payroll{}, nil
	}

	index, err := repo.getIndexByID(id)
	if err != nil {
		return &model.Payroll{}, err
	}

	payroll := payrolls[index]

	return &payroll, nil
}
