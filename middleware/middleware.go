package middleware

import (
	"log"
	"net/http"
)

type MiddlewareFunc func(http.ResponseWriter, *http.Request, http.Handler) http.Handler

type Middleware struct {
	Handler     http.Handler
	Middlewares []MiddlewareFunc
}

func (m *Middleware) Use(middlewares ...MiddlewareFunc) {
	m.Middlewares = append(m.Middlewares, middlewares...)
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Use a separate variable to hold the result of applying middlewares
	handler := m.Handler

	// Apply middlewares in order
	for _, middleware := range m.Middlewares {
		handler = middleware(w, r, handler)
	}

	// Serve the request using the final handler
	handler.ServeHTTP(w, r)
}

// Middleware 1
func Middleware1(w http.ResponseWriter, r *http.Request, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware 1: Before request")
		next.ServeHTTP(w, r)
		log.Println("Middleware 1: After request")
	})
}

// Middleware 2
func Middleware2(w http.ResponseWriter, r *http.Request, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware 2: Before request")
		next.ServeHTTP(w, r)
		log.Println("Middleware 2: After request")
	})
}

// Middleware 3
func Middleware3(w http.ResponseWriter, r *http.Request, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware 3: Before request")
		next.ServeHTTP(w, r)
		log.Println("Middleware 3: After request")
	})
}
