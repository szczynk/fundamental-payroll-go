package handler

import (
	"encoding/json"
	"fundamental-payroll-go/helper"
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/usecase"
	"net/http"
	"strconv"
	"strings"
)

type payrollHTTPHandler struct {
	PayrollUC usecase.PayrollUsecase
}

func NewPayrollHTTPHandler(
	payrollUC usecase.PayrollUsecase,
) PayrollHTTPHandler {
	return &payrollHTTPHandler{
		PayrollUC: payrollUC,
	}
}

func (handler *payrollHTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	payrolls, err := handler.PayrollUC.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&payrolls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *payrollHTTPHandler) Add(w http.ResponseWriter, r *http.Request) {
	var payrollRequest model.PayrollRequest
	err := json.NewDecoder(r.Body).Decode(&payrollRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if payrollRequest.EmployeeID <= 0 {
		http.Error(w, helper.ErrEmployeeIdNotValid, http.StatusBadRequest)
		return
	}

	if payrollRequest.TotalHariMasuk < 0 {
		http.Error(w, helper.ErrPresentDayNotValid, http.StatusBadRequest)
		return
	}

	if payrollRequest.TotalHariTidakMasuk < 0 {
		http.Error(w, helper.ErrAbsentDayNotValid, http.StatusBadRequest)
		return
	}

	payroll, err := handler.PayrollUC.Add(&payrollRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&payroll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *payrollHTTPHandler) Detail(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/payrolls/")
	if idStr == "" {
		http.Error(w, helper.ErrPayrollIdNotValid, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, helper.ErrPayrollIdNotValid, http.StatusBadRequest)
		return
	}

	if id <= 0 {
		http.Error(w, helper.ErrPayrollIdNotValid, http.StatusBadRequest)
		return
	}

	payroll, err := handler.PayrollUC.Detail(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&payroll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
