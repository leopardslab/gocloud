package bigtable

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
)

func (bigtable *Bigtable) Listtables(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://bigtableadmin.googleapis.com/v2/" + options["parent"] + "/tables"

	client := googleauth.SignJWT()

	listbigtablerequest, err := http.NewRequest("GET", url, nil)

	listbigtablerequestparam := listbigtablerequest.URL.Query()

	listbigtablerequestparam.Add("pageToken", options["pageToken"])

	if options["view"] != "" {
		listbigtablerequestparam.Add("view", options["view"])
	}

	listbigtablerequest.URL.RawQuery = listbigtablerequestparam.Encode()

	listbigtablerequest.Header.Set("Content-Type", "application/json")

	listbigtableresp, err := client.Do(listbigtablerequest)

	fmt.Println(err)

	defer listbigtableresp.Body.Close()

	body, err := ioutil.ReadAll(listbigtableresp.Body)

	listbigtableresponse := make(map[string]interface{})
	listbigtableresponse["status"] = listbigtableresp.StatusCode
	listbigtableresponse["body"] = string(body)
	resp = listbigtableresponse
	return resp, err
}

func (bigtable *Bigtable) Deletetables(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://bigtableadmin.googleapis.com/v2/" + options["name"]

	client := googleauth.SignJWT()

	Deletebigtablerequest, err := http.NewRequest("DELETE", url, nil)

	Deletebigtablerequest.Header.Set("Content-Type", "application/json")

	Deletebigtableresp, err := client.Do(Deletebigtablerequest)

	defer Deletebigtableresp.Body.Close()

	body, err := ioutil.ReadAll(Deletebigtableresp.Body)

	Deletebigtableresponse := make(map[string]interface{})
	Deletebigtableresponse["status"] = Deletebigtableresp.StatusCode
	Deletebigtableresponse["body"] = string(body)
	resp = Deletebigtableresponse
	return resp, err
}

func (bigtable *Bigtable) Describetables(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://bigtableadmin.googleapis.com/v2/" + options["name"]

	client := googleauth.SignJWT()

	Describebigtablerequest, err := http.NewRequest("GET", url, nil)

	Describebigtablerequest.Header.Set("Content-Type", "application/json")

	Describebigtableresp, err := client.Do(Describebigtablerequest)

	defer Describebigtableresp.Body.Close()

	body, err := ioutil.ReadAll(Describebigtableresp.Body)

	Describebigtableresponse := make(map[string]interface{})
	Describebigtableresponse["status"] = Describebigtableresp.StatusCode
	Describebigtableresponse["body"] = string(body)
	resp = Describebigtableresponse
	return resp, err
}

func (bigtable *Bigtable) Createtables(request interface{}) (resp interface{}, err error) {

	//POST https://bigtableadmin.googleapis.com/v2/{parent=projects/*/instances/*}/tables

	param := request.(map[string]interface{})

	var parent string

	var option Createbigtable

	for key, value := range param {
		switch key {
		case "parent":
			parentv, _ := value.(string)
			parent = parentv

		case "tableId":
			tableIdv, _ := value.(string)
			option.tableId = tableIdv

		case "table":
			tableparam, _ := value.(map[string]interface{})

			for tablekey, tablevalue := range tableparam {

				switch tablekey {

				case "name":
					namev := tablevalue.(string)
					option.table.name = namev

				case "granularity":
					granularityv := tablevalue.(string)
					option.table.granularity = granularityv

				}
			}

		case "initialSplits":
			/*
				  initialSplitsparam := make([]map[string]interface{})
					for i := 0; i < len(initialSplitsparam); i++ {
							var initialSplits InitialSplits
							for initialSplitsparam, initialSplitsparamvalue := range initialSplitsparam[i] {
								switch initialSplitsparamkey {
								    case "key":
										    initialSplits.key = initialSplitsparamvalue.(string)
								}
							}
							option.initialSplits = append(createtable.initialSplits, initialSplits)
						}
			*/
		}
	}

	fmt.Println(parent)

	Createbigtablejsonmap := make(map[string]interface{})

	Createbigtabledictnoaryconvert(option, Createbigtablejsonmap)

	Createbigtablejson, _ := json.Marshal(Createbigtablejsonmap)

	Createbigtablejsonstring := string(Createbigtablejson)

	fmt.Println(Createbigtablejsonstring)

	//	byte(`{"tableId": "bokkkya","table": { "name": "bokkkya"}}`)

	// []byte(`{"tableId": "bokkkya","table": { "name": "bokkkya"}}`)

	//	var Createbigtablejsonstringbyte = []byte(Createbigtablejsonstring)

	var Createbigtablejsonstringbyte = []byte(`{"tableId": "bokkkya","table": { }}`)

	url := "https://bigtableadmin.googleapis.com/v2/" + parent + "/tables"

	//url := https://bigtableadmin.googleapis.com/v2/{parent=projects/*}/instances

	client := googleauth.SignJWT()

	Createbigtablerequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Createbigtablejsonstringbyte))

	Createbigtablerequest.Header.Set("Content-Type", "application/json")

	Createbigtableresp, err := client.Do(Createbigtablerequest)

	defer Createbigtableresp.Body.Close()

	body, err := ioutil.ReadAll(Createbigtableresp.Body)

	Createbigtableresponse := make(map[string]interface{})
	Createbigtableresponse["status"] = Createbigtableresp.StatusCode
	Createbigtableresponse["body"] = string(body)
	resp = Createbigtableresponse
	return resp, err
}

func Createbigtabledictnoaryconvert(option Createbigtable, Createbigtablejsonmap map[string]interface{}) {

}
