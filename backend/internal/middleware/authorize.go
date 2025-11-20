package middleware

import (
	"fmt"
	"net/http"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("devince paring middleware test")
		// TODO: Extract JWT token from Authorization header (Bearer token)
		// TODO: Validate the token (check signature, expiration, etc.)
		// TODO: Extract user ID or email from token claims
		// TODO: Check if the user exists in the database
		// TODO: If validation fails, return 401 Unauthorized
		// TODO: If validation succeeds, attach user info to request context
		// TODO: Call the next handler with next.ServeHTTP(w, r)

		next.ServeHTTP(w, r)
	})
}
