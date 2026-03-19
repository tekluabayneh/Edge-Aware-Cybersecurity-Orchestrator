package main

import (
	"agent/cmd/app"
	"agent/internal/utils"
	"fmt"
	"os"
	"path/filepath"
)

//
// i am thikng to have one simple route that like fetch the token from the env and  providere to them who ever ask for th token anfroom teh valid agent we will send those waht do you thing and then tha agen won't ask everytime it will just ask if ti doe snot one it get it it willl store it in enviromane of the os and use it as many time as it want it even it it get kill weh it wake up check form os env first and if not there fetch if there is no internet ues local one which is false url and what do you thing about this staff man
//

func main() {
	folderPath, _ := utils.GetStoragePath()
	// creaet all the files
	fileTocreate := []string{"token.text", "email.text"}
	for _, v := range fileTocreate {
		Filepath := filepath.Join(folderPath, v)
		if _, err := os.Stat(Filepath); os.IsNotExist(err) {
			f, err := os.Create(Filepath)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", v, err)
			}
			f.Close()
			fmt.Printf("Successfully created: %s\n", v)
		}
	}

	_, err := utils.FetchEnv()
	if err != nil {
		fmt.Println(err)
	}

	app.App()
}
