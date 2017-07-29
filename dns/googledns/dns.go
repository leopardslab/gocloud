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

	param := request.(map[string]interface{})
	var ManagedZone, Project string
	var option Changednsrecordsets

	for key, value := range param {
		switch key {

		case "project":
			ProjectV, _ := value.(string)
			Project = ProjectV

		case "managedZone":
			ManagedZoneV, _ := value.(string)
			ManagedZone = ManagedZoneV

		case "additions":
			AdditionsV, _ := value.([]string)
			option.Additions = AdditionsV

		case "deletions":
			DeletionsV, _ := value.([]string)
			option.Deletions = DeletionsV

		case "startTime":
			StartTimeV, _ := value.(string)
			option.StartTime = StartTimeV

		case "id":
			IDV, _ := value.(string)
			option.ID = IDV

		case "kind":
			kindV, _ := value.(string)
			option.kind = kindV

		case "status":
			StatusV, _ := value.(string)
			option.Status = StatusV

		}
	}

	Changednsrecordsetsjsonmap := make(map[string]interface{})

	Changednsrecordsetsdictnoaryconvert(option, Changednsrecordsetsjsonmap)

	Changednsrecordsetsjson, _ := json.Marshal(Changednsrecordsetsjsonmap)

	Changednsrecordsetsstring := string(Changednsrecordsetsjson)

	fmt.Println(Changednsrecordsetsstring)

	var Changednsrecordsetsjsonstringbyte = []byte(Changednsrecordsetsjsonstring)

	url := "https://www.googleapis.com/dns/v1/projects/" + Project + "/managedZones" + ManagedZone + "/changes"

	client := googleauth.SignJWT()

	Changednsrecordsetsrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Changednsrecordsetsjsonstringbyte))

	Creatednsrequest.Header.Set("Content-Type", "application/json")

	Changednsrecordsetsrresp, err := client.Do(Changednsrecordsetsrequest)

	defer Changednsrecordsetsrresp.Body.Close()

	body, err := ioutil.ReadAll(Changednsrecordsetsrresp.Body)

	fmt.Println(string(body))

	return
}

func Changednsrecordsetsdictnoaryconvert(option Changednsrecordsets, Changednsrecordsetsjsonmap map[string]interface{}) {
	if len(option.Additions) != 0 {
		Changednsrecordsetsjsonmap["additions"] = option.Additions
	}

	if len(option.Deletions) != 0 {
		Changednsrecordsetsjsonmap["Deletions"] = option.Deletions
	}

	if option.ID != "" {
		Changednsrecordsetsjsonmap["id"] = option.ID
	}
	if option.Kind != "" {
		Changednsrecordsetsjsonmap["kind"] = option.Kind
	}
	if option.StartTime != "" {
		Changednsrecordsetsjsonmap["startTime"] = option.StartTime
	}
	if option.Status != "" {
		Changednsrecordsetsjsonmap["status"] = option.Status
	}
}

func (googledns *Googledns) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects" + options["project"] + "/managedZones/" + options["managedZone"] + "/changes"

	client := googleauth.SignJWT()

	ListResourcednsRecordSetsrequest, err := http.NewRequest("GET", url, nil)

	ListResourcednsRecordSetsparam := ListResourcednsRecordSetsrequest.URL.Query()

	if options["maxResults"] != "" {
		ListResourcednsRecordSetsrequestparam.Add("deviceName", options["maxResults"])
	}

	if options["pageToken"] != 0 {
		ListResourcednsRecordSetsrequestparam.Add("pageToken", options["pageToken"])
	}

	if options["sortBy"] != "" {
		ListResourcednsRecordSetsrequestparam.Add("sortBy", options["sortBy"])
	}

	if options["sortOrder"] != "" {
		ListResourcednsRecordSetsrequestparam.Add("sortOrder", options["sortOrder"])
	}

	ListResourcednsRecordSetsrequest.URL.RawQuery = ListResourcednsRecordSetsrequestparam.Encode()

	ListResourcednsRecordSetsrequest.Header.Set("Content-Type", "application/json")

	ListResourcednsRecordSetssresp, err := client.Do(ListResourcednsRecordSetsrequest)

	defer ListResourcednsRecordSetsresp.Body.Close()

	body, err := ioutil.ReadAll(ListResourcednsRecordSetsresp.Body)

	fmt.Println(string(body))

	return
}

func (googledns *Googledns) Createdns(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Project string
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
		Creatednsjsonmap["nameServers"] = option.nameServers
	}

	if option.Name != "" {
		Creatednsjsonmap["name"] = option.Name
	}

	if option.NameServerSet != "" {
		Creatednsjsonmap["NameServerSet"] = option.NameServerSet
	}

	if option.DNSName != "" {
		Creatednsjsonmap["dnsName"] = option.DNSName
	}

	if option.Kind != "" {
		Creatednsjsonmap["kind"] = option.Kind
	}

	if option.ID != "" {
		Creatednsjsonmap["id"] = option.ID
	}
	if option.Description != "" {
		Creatednsjsonmap["description"] = option.Description
	}
	if option.CreationTime != "" {
		Creatednsjsonmap["creationTime"] = option.CreationTime
	}
}
