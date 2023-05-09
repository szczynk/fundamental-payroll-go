//go:generate mockery --output=../mocks --name SalaryHandler
package handler

type SalaryHandler interface {
	List()
}
