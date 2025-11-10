package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// To Do
		// check login credential is correct and free from invalid form
		fmt.Println("check middleware work")
		next.ServeHTTP(w, r)
	})

}
