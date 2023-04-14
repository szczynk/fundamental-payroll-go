package handler

import (
	"fundamental-payroll-go/helper/logger"
	"fundamental-payroll-go/helper/response"
	"fundamental-payroll-go/usecase"
	"net/http"
)

type salaryHTTPHandler struct {
	Logger   *logger.Logger
	SalaryUC usecase.SalaryUsecase
}

func NewSalaryHTTPHandler(l *logger.Logger, salaryUC usecase.SalaryUsecase) SalaryHTTPHandler {
	return &salaryHTTPHandler{
		Logger:   l,
		SalaryUC: salaryUC,
	}
}

func (handler *salaryHTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	salaries, err := handler.SalaryUC.List()
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = response.NewJson(w, http.StatusOK, http.StatusText(http.StatusOK), salaries)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
