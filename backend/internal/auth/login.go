package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edge-aware-cyberSecurity/internal/middleware"
)

type AuthLoginHandlerType struct{}

type LoginStructure struct {
	Name  string
	Email string
}

func (h *AuthLoginHandlerType) Login(w http.ResponseWriter, r *http.Request) {
	formData, ok := r.Context().Value(middleware.LoginFromDataKey).(middleware.FormType)
	if !ok {
		log.Fatal("error loading form data")
	}

	fmt.Println(formData)
	// TODO: Check if the user exists in the database by email
	// TODO: Compare the provided password with the hashed password
	// TODO: Handle incorrect credentials safely (do not reveal which field failed)
	// TODO: Generate a JWT or session token on successful login
	// TODO: Optionally log login attempts or add rate limiting
	// TODO: Optionally add multi-factor authentication (MFA)

	fmt.Println("test work")
}
