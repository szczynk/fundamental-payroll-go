package repository

import (
	"encoding/json"
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/model"
	"os"
)

type payrollJsonRepository struct {
	jsonFile string
}

func NewPayrollJsonRepository() PayrollRepository {
	r := new(payrollJsonRepository)
	r.jsonFile = "data/payroll.json"

	return r
}

func (repo *payrollJsonRepository) encodeJSON() error {
	writer, err := os.Create(repo.jsonFile)
	if err != nil {
		return err
	}
	defer writer.Close()

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&model.Payrolls)
	if err != nil {
		return err
	}
	return nil
}

func (repo *payrollJsonRepository) decodeJSON() error {
	reader, err := os.Open(repo.jsonFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&model.Payrolls)
	if err != nil {
		return err
	}
	return nil
}

func (repo *payrollJsonRepository) List() ([]model.Payroll, error) {
	err := repo.decodeJSON()
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
		return nil, err
	}

	newPayroll := payroll
	newPayroll.ID = id + 1

	model.Payrolls = append(model.Payrolls, *newPayroll)

	err = repo.encodeJSON()
	if err != nil {
		return nil, err
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

	return -1, apperrors.New(apperrors.ErrIdNotFound)
}

func (repo *payrollJsonRepository) Detail(id int64) (*model.Payroll, error) {
	payrolls, err := repo.List()
	if err != nil {
		return nil, err
	}

	index, err := repo.getIndexByID(id)
	if err != nil {
		return nil, err
	}

	payroll := payrolls[index]

	return &payroll, nil
}
