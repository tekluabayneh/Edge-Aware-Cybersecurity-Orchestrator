package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/edge-aware-cyberSecurity/internal/utils"
)

type FormType struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Phone    string `json:"phone"`
}

func RegisterMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/json")
		var formType FormType
		if err := json.NewDecoder(r.Body).Decode(&formType); err != nil {
		}
		fmt.Println("data", formType)

		if formType.Name == "" || formType.Email == "" || formType.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "invalid form",
			})
			return
		}
		next.ServeHTTP(w, r)
	})

}

func LoginMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/json")
		var LoginFormDataType FormType

		json.NewDecoder(r.Body).Decode(&LoginFormDataType)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(body))
		if LoginFormDataType.Name == "" || LoginFormDataType.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "password or name is missing",
			})
			return
		}

		if len(LoginFormDataType.Password) < 6 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "password length must be at list 6",
			})
			return
		}

		var inputJSON []byte

		if err := json.Unmarshal(inputJSON, &LoginFormDataType); err != nil {
			fmt.Println("Error decoding input:", err)
			return
		}

		val := reflect.ValueOf(LoginFormDataType)
		for i := 0; i < val.NumField(); i++ {
			fieldValue := val.Field(i).String()

			// check if inputs are free from bad scripts
			if !utils.IsINputSafe(fieldValue) {
				json.NewEncoder(w).Encode(map[string]string{
					"message": "malformed form input or your input contain script tag",
				})
				return
			}
		}

		// check if teh input are valid and free from malware
		next.ServeHTTP(w, r)
	})
}
