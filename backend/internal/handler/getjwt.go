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

// peridodically send teh jwt toke before the previose jwt expire
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

	// otherwise store itin fiels rotate the jwt so every request will fetch those jwt
	// write one functin that fetch and reate it back to normal jwt and all those handler will fetch it and use it
	//
}
