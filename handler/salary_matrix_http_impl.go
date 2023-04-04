package handler

import (
	"encoding/json"
	"fundamental-payroll-go/usecase"
	"net/http"
)

type salaryHTTPHandler struct {
	SalaryUC usecase.SalaryUsecase
}

func NewSalaryHTTPHandler(salaryUC usecase.SalaryUsecase) SalaryHTTPHandler {
	return &salaryHTTPHandler{
		SalaryUC: salaryUC,
	}
}

func (handler *salaryHTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	salaries, err := handler.SalaryUC.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&salaries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
