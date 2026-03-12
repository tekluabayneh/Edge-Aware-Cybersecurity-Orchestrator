package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type JwtType struct {
	jwt string
}

func RequestJwt(w http.ResponseWriter, r *http.Request) {
	var jwt JwtType
	err := json.NewDecoder(r.Body).Decode(&jwt)
	if err != nil {
		fmt.Println(err)
	}

	data := jwt.jwt
	err = os.WriteFile("./jwt", []byte(data), 0644)
	if err != nil {
		fmt.Println("WRITE FILE ERROR:", err)
		return
	}
}
