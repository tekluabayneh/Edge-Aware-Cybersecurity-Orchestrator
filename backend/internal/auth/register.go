package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edge-aware-cyberSecurity/internal/middleware"
)

type AuthRegisterHandlerType struct{}

type RegisterStructure struct {
	Name     string
	Email    string
	Password string
}

func (h *AuthRegisterHandlerType) Register(w http.ResponseWriter, r *http.Request) {
	formData, ok := r.Context().Value(middleware.RegisterFormDataKey).(middleware.FormType)
	if !ok {
		log.Fatal("eror loading data")
	}
	fmt.Println("test register work", formData)

	// TODO: Check if user already exists in the database (email or phone)
	// TODO: Hash the password before saving (use bcrypt or similar)
	// TODO: Save the new user to the database
	// TODO: Optionally generate a session or JWT token for immediate login
	// TODO: Optionally send a confirmation or welcome email

}
