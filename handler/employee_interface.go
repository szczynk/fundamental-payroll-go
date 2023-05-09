//go:generate mockery --output=../mocks --name EmployeeHandler
package handler

type EmployeeHandler interface {
	List()
	Add()
}
