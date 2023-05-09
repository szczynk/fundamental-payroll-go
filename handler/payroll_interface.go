//go:generate mockery --output=../mocks --name PayrollHandler
package handler

type PayrollHandler interface {
	List()
	Add()
	Detail()
}
