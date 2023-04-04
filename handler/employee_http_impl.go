package handler

import (
	"encoding/json"
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
		http.Error(w, "Name yang dimasukkan tidak valid", http.StatusBadRequest)
		return
	}

	gender := strings.ToLower(strings.TrimSpace(employeeRequest.Gender))
	if gender == "l" {
		gender = "laki-laki"
	} else if gender == "p" {
		gender = "perempuan"
	}
	if gender != "laki-laki" && gender != "perempuan" {
		http.Error(w, "Gender yang dimasukkan tidak valid", http.StatusBadRequest)
		return
	}

	// ? are we using salaryMatrix usecase or direct hit?
	// ? currently using direct hit
	validGrades := []int8{1, 2, 3}
	var isValidGrade bool
	for _, g := range validGrades {
		if employeeRequest.Grade == g {
			isValidGrade = true
			break
		}
	}
	if !isValidGrade {
		http.Error(w, "Grade yang dimasukkan tidak valid", http.StatusBadRequest)
		return
	}

	employee, err := handler.EmployeeUC.Add(&employeeRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
