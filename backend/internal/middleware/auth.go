package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/edge-aware-cyberSecurity/internal/utils"
)

type FormType struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type contextKey string

const RegisterFormDataKey contextKey = "formDataRegister"
const LoginFromDataKey contextKey = "formDataLogin"

func RegisterMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/json")
		var RegisterFormData FormType
		if err := json.NewDecoder(r.Body).Decode(&RegisterFormData); err != nil {
		}

		if RegisterFormData.Name == "" || RegisterFormData.Email == "" || RegisterFormData.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "invalid form or missing form field",
			})
			return
		}

		if len(RegisterFormData.Password) < 6 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "password length must be at list 6",
			})
			return
		}

		if bol := checkIsValidFields(RegisterFormData); !bol {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "malformed form input or your input contain script tag",
			})
			return
		}
		ctx := context.WithValue(r.Context(), RegisterFormDataKey, RegisterFormData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func LoginMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/json")
		var LoginFormDataType FormType

		json.NewDecoder(r.Body).Decode(&LoginFormDataType)
		if LoginFormDataType.Email == "" || LoginFormDataType.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "password or email are missing",
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

		if bol := checkIsValidFields(LoginFormDataType); !bol {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "malformed form input or your input contain script tag",
			})
			return
		}
		ctx := context.WithValue(r.Context(), LoginFromDataKey, LoginFormDataType)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
func checkIsValidFields(formData FormType) bool {
	val := reflect.ValueOf(formData)
	for i := range val.NumField() {
		value := val.Field(i).String()
		if !utils.IsINputSafe(value) {
			return false
		}
	}
	return true
}
