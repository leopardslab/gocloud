package openstackauth

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func openstackauth() {

	urlv := "http://192.168.2.6/identity/v3/auth/tokens"

	var home string = os.Getenv("HOME")

	jsonFile, err := os.Open(home + "/auth.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	client := http.Client{}

	Createnoderequest, err := http.NewRequest("POST", urlv, bytes.NewBuffer(byteValue))

	if err != nil {
		fmt.Println("Createnoderequest error")
		fmt.Println(err)
	}

	Createnoderequest.Header.Set("Content-Type", "application/json")

	Createnoderesp, err := client.Do(Createnoderequest)

	if err != nil {
		fmt.Println("Createnoderesp error")
		fmt.Println(err)
	}

	defer Createnoderesp.Body.Close()

	fmt.Println("response Headers:", Createnoderesp.Header["X-Subject-Token"])
}
