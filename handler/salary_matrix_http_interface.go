package handler

import "net/http"

type SalaryHTTPHandler interface {
	List(w http.ResponseWriter, r *http.Request)
}
