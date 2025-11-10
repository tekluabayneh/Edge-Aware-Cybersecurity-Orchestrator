package handler

import (
	"fmt"
	"net/http"
)

type AuthRegisterHandlerType struct{}

type RegisterStructure struct {
	Name  string
	Email string
}

func (h *AuthRegisterHandlerType) Register(w http.ResponseWriter, r *http.Request) {

	fmt.Println("test register work")
}
