package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Print error and full stack trace
				fmt.Printf("🚨 Panic recovered: %v\n", err)
				fmt.Printf("📝 Stack trace:\n%s\n", debug.Stack())

				// Return error to client
				http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
