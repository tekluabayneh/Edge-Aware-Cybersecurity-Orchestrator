package main

import (
	"agent/cmd/app"
	"agent/internal/utils"
	"fmt"
)

//
// i am thikng to have one simple route that like fetch the token from the env and  providere to them who ever ask for th token anfroom teh valid agent we will send those waht do you thing and then tha agen won't ask everytime it will just ask if ti doe snot one it get it it willl store it in enviromane of the os and use it as many time as it want it even it it get kill weh it wake up check form os env first and if not there fetch if there is no internet ues local one which is false url and what do you thing about this staff man
//

func main() {
	_, err := utils.FetchEnv()
	if err != nil {

		fmt.Println(err)
		fmt.Println("Error loading env")
	}

	app.App()
}
