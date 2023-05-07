package middleware

import (
	"net/http"
)

func Cors(w http.ResponseWriter, r *http.Request, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
