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

	CreateNoderequest, err := http.NewRequest("POST", urlv, bytes.NewBuffer(byteValue))

	if err != nil {
		fmt.Println("CreateNoderequest error")
		fmt.Println(err)
	}

	CreateNoderequest.Header.Set("Content-Type", "application/json")

	CreateNoderesp, err := client.Do(CreateNoderequest)

	if err != nil {
		fmt.Println("CreateNoderesp error")
		fmt.Println(err)
	}

	defer CreateNoderesp.Body.Close()

	fmt.Println("response Headers:", CreateNoderesp.Header["X-Subject-Token"])
}
