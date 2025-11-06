package handler

import (
	"fmt"
	"net/http"
)

type AuthLoginHandlerType struct{}

type LoginStructure struct {
	Name  string
	Email string
}

func (h *AuthLoginHandlerType) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test work")
}
