package main

import (
	"database/sql"
	"fundamental-payroll-go/config"
	"fundamental-payroll-go/config/db"
	"fundamental-payroll-go/handler"
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/helper/logger"
	"fundamental-payroll-go/middleware"
	"fundamental-payroll-go/repository"
	"fundamental-payroll-go/usecase"
	"log"
	"net/http"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	l := logger.New(config.Debug)

	employeeUC, salaryUC, payrollUC := createUsecase(config, l)

	switch config.Handler {
	case "http":
		employeeHTTPHandler := handler.NewEmployeeHTTPHandler(l, employeeUC)
		salaryHTTPHandler := handler.NewSalaryHTTPHandler(l, salaryUC)
		payrollHTTPHandler := handler.NewPayrollHTTPHandler(l, payrollUC)
		err := NewServer(config, l, employeeHTTPHandler, payrollHTTPHandler, salaryHTTPHandler)
		if err != nil {
			l.Fatal().Err(err).Msg("server fail to start")
		}
	default:
		employeeHandler := handler.NewEmployeeHandler(employeeUC)
		salaryHandler := handler.NewSalaryHandler(salaryUC)
		payrollHandler := handler.NewPayrollHandler(payrollUC)
		handler.Menu(employeeHandler, payrollHandler, salaryHandler)
	}
}

func createUsecase(config *config.Config, logger *logger.Logger) (
	usecase.EmployeeUsecase,
	usecase.SalaryUsecase,
	usecase.PayrollUsecase,
) {
	var (
		DB           *sql.DB
		employeeRepo repository.EmployeeRepository
		salaryRepo   repository.SalaryRepository
		payrollRepo  repository.PayrollRepository
		err          error
	)

	switch config.Storage {
	case "sql":
		switch config.Database.Driver {
		case "pgx":
			DB, err = db.NewPgxDatabase(config)
			if err != nil {
				logger.Fatal().Err(err).Msg("database fail to start")
			}
			employeeRepo = repository.NewEmployeePgxRepository(DB)
			salaryRepo = repository.NewSalaryPgxRepository(DB)
			payrollRepo = repository.NewPayrollPgxRepository(DB)
		default:
			logger.Fatal().Msg(apperrors.ErrDbDriverNotFound)
		}
	case "json":
		employeeRepo = repository.NewEmployeeJsonRepository()
		payrollRepo = repository.NewPayrollJsonRepository()
		salaryRepo = repository.NewSalaryJsonRepository()
	default:
		employeeRepo = repository.NewEmployeeRepository()
		payrollRepo = repository.NewPayrollRepository()
		salaryRepo = repository.NewSalaryRepository()
	}
	return usecase.NewEmployeeUsecase(employeeRepo, salaryRepo),
		usecase.NewSalaryUsecase(salaryRepo),
		usecase.NewPayrollUsecase(employeeRepo, payrollRepo, salaryRepo)
}

func NewServer(
	config *config.Config,
	logger *logger.Logger,
	employeeHandler handler.EmployeeHTTPHandler,
	payrollHandler handler.PayrollHTTPHandler,
	salaryHandler handler.SalaryHTTPHandler,
) error {
	mux := http.NewServeMux()

	muxMiddleware := new(middleware.Middleware)
	muxMiddleware.Handler = mux

	muxMiddleware.Use(middleware.ContentTypeJson)
	muxMiddleware.Use(
		func(w http.ResponseWriter, r *http.Request, next http.Handler) http.Handler {
			return middleware.Log(logger, w, r, next)
		},
	)
	muxMiddleware.Use(
		func(w http.ResponseWriter, r *http.Request, next http.Handler) http.Handler {
			return middleware.Error(logger, w, r, next)
		},
	)

	mux.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			employeeHandler.List(w, r)
		case "POST":
			employeeHandler.Add(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	})

	mux.HandleFunc("/payrolls", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			payrollHandler.List(w, r)
		case "POST":
			payrollHandler.Add(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	})

	mux.HandleFunc("/payrolls/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			payrollHandler.Detail(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	})

	mux.HandleFunc("/salaries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			salaryHandler.List(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	})

	server := new(http.Server)
	server.Addr = "localhost:" + config.Port
	server.Handler = muxMiddleware

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	logger.Info().Msgf("live on http://localhost:%s", config.Port)
	return nil
}
