package handler

import (
	"encoding/json"
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/helper/logger"
	"fundamental-payroll-go/helper/response"
	"fundamental-payroll-go/model"
	"fundamental-payroll-go/usecase"
	"net/http"
	"strings"
)

type employeeHTTPHandler struct {
	Logger     *logger.Logger
	EmployeeUC usecase.EmployeeUsecase
}

func NewEmployeeHTTPHandler(l *logger.Logger, employeeUC usecase.EmployeeUsecase) EmployeeHTTPHandler {
	return &employeeHTTPHandler{
		Logger:     l,
		EmployeeUC: employeeUC,
	}
}

func (handler *employeeHTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	employees, err := handler.EmployeeUC.List()
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		_ = response.NewJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = response.NewJson(w, http.StatusOK, http.StatusText(http.StatusOK), employees)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		_ = response.NewJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (handler *employeeHTTPHandler) Add(w http.ResponseWriter, r *http.Request) {
	var employeeRequest model.EmployeeRequest
	err := json.NewDecoder(r.Body).Decode(&employeeRequest)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		_ = response.NewJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if employeeRequest.Name == "" {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrEmployeeNameNotValid, nil)
		return
	}

	gender := strings.TrimSpace(employeeRequest.Gender)
	if gender == "" {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrEmployeeGenderNotValid, nil)
		return
	}

	lowerGender := strings.ToLower(gender)
	switch lowerGender {
	case "l", "laki-laki":
		gender = "laki-laki"
	case "p", "perempuan":
		gender = "perempuan"
	default:
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrEmployeeGenderNotValid, nil)
		return
	}
	employeeRequest.Gender = gender

	if employeeRequest.Grade <= 0 {
		_ = response.NewJson(w, http.StatusBadRequest, apperrors.ErrEmployeeGradeNotValid, nil)
		return
	}

	employee, err := handler.EmployeeUC.Add(&employeeRequest)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		code, message := apperrors.HandleAppError(err)
		_ = response.NewJson(w, code, message, nil)
		return
	}

	err = response.NewJson(w, http.StatusCreated, http.StatusText(http.StatusCreated), employee)
	if err != nil {
		handler.Logger.Error().Err(err).Msg("")
		_ = response.NewJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}
