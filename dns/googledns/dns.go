package googledns

//Googledns struct represents Googledns attribute and methods associates with it.
type Googledns struct {
}

//Createdns struct represents Createdns attribute.
type Createdns struct {
	CreationTime  string   `json:"creationTime"`
	Description   string   `json:"description"`
	DNSName       string   `json:"dnsName"`
	ID            string   `json:"id"`
	Kind          string   `json:"kind"`
	Name          string   `json:"name"`
	NameServerSet string   `json:"nameServerSet"`
	NameServers   []string `json:"nameServers"`
}

type Changednsrecordsets struct {
	Additions []interface{} `json:"additions"`
	Deletions []interface{} `json:"deletions"`
	ID        string        `json:"id"`
	Kind      string        `json:"kind"`
	StartTime string        `json:"startTime"`
	Status    string        `json:"status"`
}

func (googledns *Googledns) Changednsrecordsets(request interface{}) (resp interface{}, err error) {

}

func (googledns *Googledns) Createdns(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var ManagedZone, Project string
	var option Createdns

	for key, value := range param {
		switch key {

		case "Project":
			ProjectV, _ := value.(string)
			Project = ProjectV

		case "CreationTime":
			CreationTimeV, _ := value.(string)
			option.creationTime = CreationTimeV

		case "Description":
			DescriptionV, _ := value.(string)
			option.description = DescriptionV

		case "dnsName":
			DNSNameV, _ := value.(string)
			option.DNSName = DNSNameV

		case "nameServers":
			nameServersV, _ := value.([]string)
			option.nameServers = nameServersV

		case "id":
			IDV, _ := value.(string)
			option.ID = IDV

		case "Kind":
			kindV, _ := value.(string)
			option.kind = kindV

		case "Name":
			NameV, _ := value.(string)
			option.Name = NameV

		case "nameServerSet":
			NameServerSetV, _ := value.(string)
			option.NameServerSet = NameServerSetV
		}
	}

	Creatednsjsonmap := make(map[string]interface{})

	Creatednsedictnoaryconvert(option, Creatednsjsonmap)

	Creatednsjson, _ := json.Marshal(Creatednsjsonmap)

	Creatednsjsonstring := string(Creatednsjson)

	fmt.Println(Creatednsjsonstring)

	var Creatednsjsonstringbyte = []byte(Creatednsjsonstring)

	url := "https://www.googleapis.com/dns/v1/projects/" + Project + "/managedZones"

	client := googleauth.SignJWT()

	Creatednsrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Creatednsjsonstringbyte))

	Creatednsrequest.Header.Set("Content-Type", "application/json")

	Creatednsrresp, err := client.Do(Creatednsrequest)

	defer Creatednsrresp.Body.Close()

	body, err := ioutil.ReadAll(Creatednsrresp.Body)

	fmt.Println(string(body))

	return

}

//Listdns lists DNS.
func (googledns *Googledns) Listdns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["project"] + "/managedZones/"

	client := googleauth.SignJWT()

	Listdnsrequest, err := http.NewRequest("GET", url, nil)

	Listdnsrequestparam := Listdnsrequest.URL.Query()

	if options["maxResults"] != "" {
		Listdnsrequestparam.Add("deviceName", options["maxResults"])
	}

	if options["pageToken"] != 0 {
		Listdnsrequestparam.Add("pageToken", options["pageToken"])
	}

	Listdnsrequest.URL.RawQuery = Listdnsrequestparam.Encode()

	Listdnsrequest.Header.Set("Content-Type", "application/json")

	Listdnssresp, err := client.Do(Listdnsrequest)

	defer Listdnsresp.Body.Close()

	body, err := ioutil.ReadAll(Listdnsresp.Body)

	fmt.Println(string(body))

	return
}

//Deletedns deletes DNS.
func (googledns *Googledns) Deletedns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["project"] + "/zones/" + options["Zone"] + "/operations/" + options["operation"]

	client := googleauth.SignJWT()

	Deletednsrequest, err := http.NewRequest("DELETE", url, nil)

	Deletednsrequest.Header.Set("Content-Type", "application/json")

	Deletednsresp, err := client.Do(Deletednsrequest)

	defer Deletednsresp.Body.Close()

	body, err := ioutil.ReadAll(Deletednsresp.Body)

	fmt.Println(string(body))

	return
}

func Creatednsedictnoaryconvert(option Createdns, Creatednsjsonmap map[string]interface{}) {

	if len(option.nameServers) != 0 {
		Createclusterjsonmap["nameServers"] = option.nameServers
	}

	if option.Name != "" {
		Createclusterjsonmap["name"] = option.Name
	}

	if option.NameServerSet != "" {
		Createclusterjsonmap["NameServerSet"] = option.NameServerSet
	}

	if option.DNSName != "" {
		Createclusterjsonmap["dnsName"] = option.DNSName
	}

	if option.Kind != "" {
		Createclusterjsonmap["kind"] = option.Kind
	}

	if option.ID != "" {
		Createclusterjsonmap["id"] = option.ID
	}
	if option.Description != "" {
		Createclusterjsonmap["description"] = option.Description
	}
	if option.CreationTime != "" {
		Createclusterjsonmap["creationTime"] = option.CreationTime
	}
}
