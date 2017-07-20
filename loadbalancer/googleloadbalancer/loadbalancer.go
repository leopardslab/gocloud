package googleloadbalancer

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/scorelab/gocloud-v2/googleauth"
	"io/ioutil"
	"net/http"
)


type TargetPools struct {
	BackupPool        string   `json:"backupPool"`
	CreationTimestamp string   `json:"creationTimestamp"`
	Description       string   `json:"description"`
	HealthChecks      []string `json:"healthChecks"`
	FailoverRatio     int      `json:"failoverRatio"`
	ID                string   `json:"id"`
	Instances         []string `json:"instances"`
	Kind              string   `json:"kind"`
	Name              string   `json:"name"`
	Region            string   `json:"region"`
	SelfLink          string   `json:"selfLink"`
	SessionAffinity   string   `json:"sessionAffinity"`
}

func (googleloadbalancer *Googleloadbalancer) Creatloadbalancer(request interface{}) (resp interface{}, err error) {

	var option TargetPools

	var Project string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "Project":
			Project, _ = value.(string)

		case "name":
			name, _ := value.(string)
			option.Name = name

		case "Region":
			RegionV, _ := value.(string)
			option.Region = RegionV

		case "healthChecks":
			HealthChecksV, _ := value.([]string)
			option.HealthChecks = HealthChecksV

		case "description":
			DescriptionV, _ := value.(string)
			option.Description = DescriptionV

		case "BackupPool":
			BackupPoolV, _ := value.(string)
			option.BackupPool = BackupPoolV

		case "failoverRatio":
			FailoverRatioV, _ := value.(int)
			option.FailoverRatio = FailoverRatioV

		case "id":
			IDV, _ := value.(string)
			option.ID = IDV

		case "instances":
			InstancesV, _ := value.([]string)
			option.Instances = InstancesV

		case "kind":
			KindV, _ := value.(string)
			option.Kind = KindV

		case "sessionAffinity":
			SessionAffinityV, _ := value.(string)
			option.SessionAffinity = SessionAffinityV

		case "region":
			RegionV, _ := value.(string)
			option.Region = RegionV

		case "selfLink":
			SelfLinkV, _ := value.(string)
			option.SelfLink = SelfLinkV

		}
	}

	option.CreationTimestamp = ""

	Creatloadbalancerjsonmap := make(map[string]interface{})

	Creatloadbalancerdictnoaryconvert(option, Creatloadbalancerjsonmap)

	Creatloadbalancerjson, _ := json.Marshal(Creatloadbalancerjsonmap)

	Creatloadbalancerjsonstring := string(Creatloadbalancerjson)

	fmt.Println(Creatloadbalancerjsonstring)

	var Creatloadbalancerstringbyte = []byte(Creatloadbalancerjsonstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + Project + "regions/" + option.Region + "targetPools"

	client := googleauth.SignJWT()

	Creatloadbalancerrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Creatloadbalancerstringbyte))

	Creatloadbalancerrequest.Header.Set("Content-Type", "application/json")

	Creatloadbalancerresp, err := client.Do(Creatloadbalancerrequest)

	defer Creatloadbalancerresp.Body.Close()

	body, err := ioutil.ReadAll(Creatloadbalancerresp.Body)

	fmt.Println(string(body))

	return

}

func Creatloadbalancerdictnoaryconvert(option TargetPools, Creatloadbalancerjsonmap map[string]interface{}){

	if(option.ID == ""){
		Creatloadbalancerjsonmap["id"]= option.ID
	}

	if(len(option.BackupPool) != 0){
		Creatloadbalancerjsonmap["backupPool"]= option.BackupPool
	}

	if(len(option.Instances) != 0){
		Creatloadbalancerjsonmap["instances"]= option.Instances
	}

	if(option.FailoverRatio != 0){
		Creatloadbalancerjsonmap["failoverRatio"]= option.FailoverRatio
	}

	if(option.CreationTimestamp == ""){
		Creatloadbalancerjsonmap["creationTimestamp"]= option.CreationTimestamp
	}

	if(option.Description == ""){
		Creatloadbalancerjsonmap["description"]= option.Description
	}

	if(len(option.HealthChecks) != 0){
		Creatloadbalancerjsonmap["healthChecks"]= option.HealthChecks
	}

	if(option.SessionAffinity == ""){
		Creatloadbalancerjsonmap["sessionAffinity"]= option.SessionAffinity
	}

	if(option.SelfLink == ""){
		Creatloadbalancerjsonmap["selfLink"]= option.SelfLink
	}
	if(option.Region == ""){
		Creatloadbalancerjsonmap["region"]= option.Region
	}

	if(option.Name == ""){
		Creatloadbalancerjsonmap["name"]= option.Name
	}

	if(option.Kind == ""){
		Creatloadbalancerjsonmap["kind"]= option.Kind
	}

}


func (googleloadbalancer *Googleloadbalancer) Deleteloadbalancer(request interface{}) (resp interface{}, err error) {

		options := request.(map[string]string)

		url := "https://www.googleapis.com/compute/v1/projects/" + options["project"] + "/regions/" + options["region"] + "/targetPools/" + options["targetPool"]

		client := googleauth.SignJWT()

		Deleteloadbalancerrequest, err := http.NewRequest("DELETE", url, nil)
		Deleteloadbalancerrequest.Header.Set("Content-Type", "application/json")

		Deleteloadbalancerresp, err := client.Do(Deleteloadbalancerrequest)

		defer Deleteloadbalancerresp.Body.Close()

		body, err := ioutil.ReadAll(Deleteloadbalancerresp.Body)

		fmt.Println(string(body))

		return
}

func (googleloadbalancer *Googleloadbalancer) Listloadbalancer(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["project"] + "/regions/" + options["region"] + "/targetPools/"

	client := googleauth.SignJWT()

	Listloadbalancerrequest, err := http.NewRequest("GET", url, nil)
	Listloadbalancerrequest.Header.Set("Content-Type", "application/json")

	Listloadbalancerresp, err := client.Do(Listloadbalancerrequest)

	defer Listloadbalancerresp.Body.Close()

	body, err := ioutil.ReadAll(Listloadbalancerresp.Body)

	fmt.Println(string(body))

	return
}


func (googleloadbalancer *Googleloadbalancer) Attachnodewithloadbalancer(request interface{}) (resp interface{}, err error) {

	var TargetPool string

	var Project string

	var Region string

	var Instances []string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "Project":
			ProjectV, _ := value.(string)
			Project = ProjectV

		case "TargetPool":
			TargetPoolV, _ := value.(string)
			TargetPool = TargetPoolV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "Instances":
			InstancesV, _ := value.([]string)
			Instances = InstancesV
		}
}
	url := "https://www.googleapis.com/compute/v1/projects/" + Project + "/regions/" + Region + "/targetPools/" + TargetPool +"/addInstance"

	Attachnodewithloadbalancerjsonmap := make(map[string]interface{})


	if(len(Instances)!=0){
		instance := []interface{}{}
		for i :=0;i<(len(Instances));i++{
				val := map[string]string{}
				val["instance"] = Instances[i]
				instance = append(instance,val)
		}
		Attachnodewithloadbalancerjsonmap["instances"] = instance
	}

	Attachnodewithloadbalancerjson, _ := json.Marshal(Attachnodewithloadbalancerjsonmap)

	Attachnodewithloadbalancerjsonstring := string(Attachnodewithloadbalancerjson)

	fmt.Println(Attachnodewithloadbalancerjsonstring)

	var Attachnodewithloadbalancerstringbyte = []byte(Attachnodewithloadbalancerjsonstring)


	client := googleauth.SignJWT()

	Attachnodewithloadbalancerrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Attachnodewithloadbalancerstringbyte))

	Attachnodewithloadbalancerrequest.Header.Set("Content-Type", "application/json")

	Attachnodewithloadbalancerresp, err := client.Do(Attachnodewithloadbalancerrequest)

	defer Attachnodewithloadbalancerresp.Body.Close()

	body, err := ioutil.ReadAll(Attachnodewithloadbalancerresp.Body)

	fmt.Println(string(body))

	return
}


func (googleloadbalancer *Googleloadbalancer) Detachnodewithloadbalancer(request interface{}) (resp interface{}, err error) {
	var TargetPool string

	var Project string

	var Region string

	var Instances []string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "Project":
			ProjectV, _ := value.(string)
			Project = ProjectV

		case "TargetPool":
			TargetPoolV, _ := value.(string)
			TargetPool = TargetPoolV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "Instances":
			InstancesV, _ := value.([]string)
			Instances = InstancesV
		}
}
	url := "https://www.googleapis.com/compute/v1/projects/" + Project + "/regions/" + Region + "/targetPools/" + TargetPool +"/addInstance"

	Detachnodewithloadbalancerjsonmap := make(map[string]interface{})


	if(len(Instances)!=0){
		instance := []interface{}{}
		for i :=0;i<(len(Instances));i++{
				val := map[string]string{}
				val["instance"] = Instances[i]
				instance = append(instance,val)
		}
		Detachnodewithloadbalancerjsonmap["instances"] = instance
	}

	Detachnodewithloadbalancerjson, _ := json.Marshal(Detachnodewithloadbalancerjsonmap)

	Detachnodewithloadbalancerjsonstring := string(Detachnodewithloadbalancerjson)

	fmt.Println(Detachnodewithloadbalancerjsonstring)

	var Detachnodewithloadbalancerstringbyte = []byte(Detachnodewithloadbalancerjsonstring)

	client := googleauth.SignJWT()

	Detachnodewithloadbalancerrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Detachnodewithloadbalancerstringbyte))

	Detachnodewithloadbalancerrequest.Header.Set("Content-Type", "application/json")

	Detachnodewithloadbalancerresp, err := client.Do(Detachnodewithloadbalancerrequest)

	defer Detachnodewithloadbalancerresp.Body.Close()

	body, err := ioutil.ReadAll(Detachnodewithloadbalancerresp.Body)

	fmt.Println(string(body))

	return

}
