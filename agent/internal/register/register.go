package register

import (
	"errors"
	"fmt"
)

func Register() (bool, error) {
	fmt.Println("register use ")
	if 1 > 3 {
		return true, nil
	}
	return false, errors.New("user not register")
}
