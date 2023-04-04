package handler

import "net/http"

type EmployeeHTTPHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
}
