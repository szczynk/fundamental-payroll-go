package middleware

import (
	"fundamental-payroll-go/helper/logger"
	"fundamental-payroll-go/helper/response"
	"net/http"
)

func Error(logger *logger.Logger, w http.ResponseWriter, r *http.Request, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				// Log the server error
				logger.Error().Msgf("Server error: %v", r)
				_ = response.NewJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
