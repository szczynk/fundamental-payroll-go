package handler

import (
	"encoding/json"
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/helper/logger"
	"fundamental-payroll-go/helper/response"
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/usecase"
	"net/http"
	"strconv"
	"strings"
)

type payrollHTTPHandler struct {
	Logger    *logger.Logger
	PayrollUC usecase.PayrollUsecase
}

func NewPayrollHTTPHandler(l *logger.Logger, payrollUC usecase.PayrollUsecase) PayrollHTTPHandler {
	return &payrollHTTPHandler{
		Logger:    l,
		PayrollUC: payrollUC,
	}
}

func (handler *payrollHTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	payrolls, err := handler.PayrollUC.List()
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		_ = response.NewJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = response.NewJson(w, http.StatusOK, http.StatusText(http.StatusOK), payrolls)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		_ = response.NewJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (handler *payrollHTTPHandler) Add(w http.ResponseWriter, r *http.Request) {
	var payrollRequest model.PayrollRequest
	err := json.NewDecoder(r.Body).Decode(&payrollRequest)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		_ = response.NewJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if payrollRequest.EmployeeID <= 0 {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrEmployeeIdNotValid, nil)
		return
	}

	if payrollRequest.TotalHariMasuk < 0 {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrPresentDayNotValid, nil)
		return
	}

	if payrollRequest.TotalHariTidakMasuk < 0 {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrAbsentDayNotValid, nil)
		return
	}

	payroll, err := handler.PayrollUC.Add(&payrollRequest)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		code, message := apperrors.HandleAppError(err)
		_ = response.NewJson(w, code, message, nil)
		return
	}

	err = response.NewJson(w, http.StatusCreated, http.StatusText(http.StatusCreated), payroll)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		_ = response.NewJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (handler *payrollHTTPHandler) Detail(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrPayrollIdNotValid, nil)
		return
	}

	idStr := parts[2]
	if idStr == "" {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrPayrollIdNotValid, nil)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrPayrollIdNotValid, nil)
		return
	}

	if id <= 0 {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrPayrollIdNotValid, nil)
		return
	}

	payroll, err := handler.PayrollUC.Detail(id)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		code, message := apperrors.HandleAppError(err)
		_ = response.NewJson(w, code, message, nil)
		return
	}

	err = response.NewJson(w, http.StatusOK, http.StatusText(http.StatusOK), payroll)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		_ = response.NewJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}
