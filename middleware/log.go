package middleware

import (
	"fundamental-payroll-go/helper/logger"
	"net/http"
	"time"
)

func Log(logger *logger.Logger, w http.ResponseWriter, r *http.Request, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log request
		start := time.Now()

		next.ServeHTTP(w, r)

		latency := time.Since(start)

		logger.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Str("user_agent", r.UserAgent()).
			Str("referer", r.Referer()).
			Str("proto", r.Proto).
			Str("remote_ip", r.RemoteAddr).
			Dur("latency", latency).
			Msg("")
	})
}
