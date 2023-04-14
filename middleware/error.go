package middleware

import (
	"fundamental-payroll-go/helper/logger"
	"net/http"
)

func Error(logger *logger.Logger, w http.ResponseWriter, r *http.Request, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				// Log the server error
				logger.Error().Msgf("Server error: %v", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
