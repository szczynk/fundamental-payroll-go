package handler

import (
	"encoding/json"
	"fundamental-payroll-go/helper"
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/usecase"
	"net/http"
	"strings"
)

type employeeHTTPHandler struct {
	EmployeeUC usecase.EmployeeUsecase
}

func NewEmployeeHTTPHandler(employeeUC usecase.EmployeeUsecase) EmployeeHTTPHandler {
	return &employeeHTTPHandler{
		EmployeeUC: employeeUC,
	}
}

func (handler *employeeHTTPHandler) List(w http.ResponseWriter, r *http.Request) {

	employees, err := handler.EmployeeUC.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *employeeHTTPHandler) Add(w http.ResponseWriter, r *http.Request) {
	var employeeRequest model.EmployeeRequest
	err := json.NewDecoder(r.Body).Decode(&employeeRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if employeeRequest.Name == "" {
		http.Error(w, helper.ErrEmployeeNameNotValid, http.StatusBadRequest)
		return
	}

	gender := strings.ToLower(strings.TrimSpace(employeeRequest.Gender))
	if gender == "l" {
		gender = "laki-laki"
	} else if gender == "p" {
		gender = "perempuan"
	}
	if gender != "laki-laki" && gender != "perempuan" {
		http.Error(w, helper.ErrEmployeeGenderNotValid, http.StatusBadRequest)
		return
	}
	employeeRequest.Gender = gender

	if employeeRequest.Grade <= 0 {
		http.Error(w, helper.ErrEmployeeGradeNotValid, http.StatusBadRequest)
		return
	}

	employee, err := handler.EmployeeUC.Add(&employeeRequest)
	if err != nil {
		if appErr, ok := err.(*helper.AppError); ok {
			http.Error(w, appErr.Message, http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
