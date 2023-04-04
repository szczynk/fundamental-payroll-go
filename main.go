package main

import (
	"fundamental-payroll-go/config"
	"fundamental-payroll-go/config/db"
	"fundamental-payroll-go/handler"
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

	employeeUC, salaryUC, payrollUC := createUsecase(config)

	switch config.Mode {
	case "http":
		employeeHTTPHandler := handler.NewEmployeeHTTPHandler(employeeUC)
		salaryHTTPHandler := handler.NewSalaryHTTPHandler(salaryUC)
		payrollHTTPHandler := handler.NewPayrollHTTPHandler(payrollUC)
		err := NewServer(config, employeeHTTPHandler, payrollHTTPHandler, salaryHTTPHandler)
		if err != nil {
			log.Fatal(err)
		}
	default:
		employeeHandler := handler.NewEmployeeHandler(employeeUC)
		salaryHandler := handler.NewSalaryHandler(salaryUC)
		payrollHandler := handler.NewPayrollHandler(payrollUC)
		handler.Menu(employeeHandler, payrollHandler, salaryHandler)
	}
}

func createUsecase(config *config.Config) (
	usecase.EmployeeUsecase,
	usecase.SalaryUsecase,
	usecase.PayrollUsecase,
) {
	var (
		employeeRepo repository.EmployeeRepository
		salaryRepo   repository.SalaryRepository
		payrollRepo  repository.PayrollRepository
	)

	switch config.Storage {
	case "sql":
		switch config.Database.Driver {
		case "pgx":
			db, err := db.NewPgxDatabase(config)
			if err != nil {
				log.Fatal(err)
			}
			employeeRepo = repository.NewEmployeePgxRepository(db)
			salaryRepo = repository.NewSalaryPgxRepository(db)
			payrollRepo = repository.NewPayrollPgxRepository(db)
		default:
			log.Fatalln("database driver not existed")
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
	return usecase.NewEmployeeUsecase(employeeRepo),
		usecase.NewSalaryUsecase(salaryRepo),
		usecase.NewPayrollUsecase(employeeRepo, payrollRepo, salaryRepo)
}

func NewServer(
	config *config.Config,
	employeeHandler handler.EmployeeHTTPHandler,
	payrollHandler handler.PayrollHTTPHandler,
	salaryHandler handler.SalaryHTTPHandler,
) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.Method {
		case "GET":
			employeeHandler.List(w, r)
		case "POST":
			employeeHandler.Add(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	mux.HandleFunc("/payrolls", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.Method {
		case "GET":
			payrollHandler.List(w, r)
		case "POST":
			payrollHandler.Add(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	mux.HandleFunc("/payrolls/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.Method {
		case "GET":
			payrollHandler.Detail(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	mux.HandleFunc("/salaries", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.Method {
		case "GET":
			salaryHandler.List(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	server := &http.Server{
		Addr:    "localhost:" + config.Port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	log.Printf("live on http://localhost:%s", config.Port)
	return nil
}
