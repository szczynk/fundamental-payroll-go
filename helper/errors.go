package helper

const (
	ErrPlatformNotSupported  = "your platform is unsupported! i can't clear terminal screen :("
	ErrIdNotFound            = "id not found"
	ErrDbUrlNotExist         = "database URL not found"
	ErrEnvNotFound           = ".env file not found"
	ErrEmployeeGradeNotValid = "invalid employee grade"

	ErrMethodNotAllowed       = "method not allowed"
	ErrEmployeeNameNotValid   = "name yang dimasukkan tidak valid"
	ErrEmployeeGenderNotValid = "gender yang dimasukkan tidak valid"
	ErrPayrollIdNotValid      = "payroll id yang dimasukkan tidak valid"
	ErrEmployeeIdNotValid     = "employee id yang dimasukkan tidak valid"
	ErrPresentDayNotValid     = "total hari masuk yang dimasukkan tidak valid"
	ErrAbsentDayNotValid      = "total hari tidak masuk yang dimasukkan tidak valid"
)

type AppError struct {
	Message string
}

func NewAppError(message string) *AppError {
	return &AppError{
		Message: message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}
