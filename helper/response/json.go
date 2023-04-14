package response

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewJson(w http.ResponseWriter, code int, message string, data interface{}) error {
	res := jsonResponse{
		Status:  code,
		Message: message,
		Data:    data,
	}

	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(&res)
}
