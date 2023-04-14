package apperrors

import "net/http"

const (
	ErrPlatformNotSupported = "your platform is unsupported! i can't clear terminal screen :("
	ErrIdNotFound           = "id not found"
	ErrDbUrlNotFound        = "database url not found"
	ErrDbDriverNotFound     = "database driver not found"
	ErrEnvNotFound          = ".env file not found"

	ErrEmployeeIdNotValid     = "employee id yang dimasukkan tidak valid"
	ErrEmployeeGradeNotValid  = "invalid employee grade"
	ErrEmployeeNameNotValid   = "name yang dimasukkan tidak valid"
	ErrEmployeeGenderNotValid = "gender yang dimasukkan tidak valid"
	ErrEmployeeNotFound       = "employee not found"

	ErrPayrollIdNotValid  = "payroll id yang dimasukkan tidak valid"
	ErrPresentDayNotValid = "total hari masuk yang dimasukkan tidak valid"
	ErrAbsentDayNotValid  = "total hari tidak masuk yang dimasukkan tidak valid"
	ErrPayrollNotFound    = "payroll not found"
)

func HandleAppError(err error) (int, string) {
	switch e := err.(type) {
	case *AppError:
		switch e.Message {
		case ErrEmployeeNotFound:
			return http.StatusNotFound, err.Error()
		case ErrPayrollNotFound:
			return http.StatusNotFound, err.Error()
		default:
			return http.StatusInternalServerError, err.Error()
		}

	default:
		return http.StatusInternalServerError, err.Error()
	}
}

type AppError struct {
	Message string
}

func New(message string) *AppError {
	return &AppError{
		Message: message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}
