package googleloadbalancer

import (
	"bytes"
	"encoding/json"
	googleauth "github.com/scorelab/gocloud-v2/googleauth"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	RFC3339  = "2006-01-02T15:04:05Z07:00"
)

//Creatloadbalancer creates google loadbalancer pool.
func (googleloadbalancer *Googleloadbalancer) Creatloadbalancer(request interface{}) (resp interface{}, err error) {

	var option TargetPools
	var Project string
	var Region string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "Project":
			Project, _ = value.(string)

		case "Name":
			name, _ := value.(string)
			option.Name = name

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

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

		case "Instances":
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
			Region = RegionV

		case "selfLink":
			SelfLinkV, _ := value.(string)
			option.SelfLink = SelfLinkV

		}
	}

	option.Region = "https://www.googleapis.com/compute/v1/projects/" + Project + "zones/" + Region

	option.CreationTimestamp = time.Now().UTC().Format(time.RFC3339)


	Creatloadbalancerjsonmap := make(map[string]interface{})

	Creatloadbalancerdictnoaryconvert(option, Creatloadbalancerjsonmap)

	Creatloadbalancerjson, _ := json.Marshal(Creatloadbalancerjsonmap)

	Creatloadbalancerjsonstring := string(Creatloadbalancerjson)

	var Creatloadbalancerstringbyte = []byte(Creatloadbalancerjsonstring)

	url := "https://www.googleapis.com/compute/beta/projects/" + Project + "/regions/" + Region + "/targetPools"

	client := googleauth.SignJWT()

	Creatloadbalancerrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Creatloadbalancerstringbyte))

	Creatloadbalancerrequest.Header.Set("Content-Type", "application/json")

	Creatloadbalancerresp, err := client.Do(Creatloadbalancerrequest)

	defer Creatloadbalancerresp.Body.Close()

	body, err := ioutil.ReadAll(Creatloadbalancerresp.Body)

	Creatloadbalancerresponse := make(map[string]interface{})
	Creatloadbalancerresponse["status"] = Creatloadbalancerresp.StatusCode
	Creatloadbalancerresponse["body"] = string(body)
	resp = Creatloadbalancerresponse
	return resp, err
}

//Deleteloadbalancer deletes google loadbalancer pool.
func (googleloadbalancer *Googleloadbalancer) Deleteloadbalancer(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/beta/projects/" + options["Project"] + "/regions/" + options["Region"] + "/targetPools/" + options["TargetPool"]

	client := googleauth.SignJWT()

	Deleteloadbalancerrequest, err := http.NewRequest("DELETE", url, nil)
	Deleteloadbalancerrequest.Header.Set("Content-Type", "application/json")

	Deleteloadbalancerresp, err := client.Do(Deleteloadbalancerrequest)

	defer Deleteloadbalancerresp.Body.Close()

	body, err := ioutil.ReadAll(Deleteloadbalancerresp.Body)

	Deleteloadbalancerresponse := make(map[string]interface{})
	Deleteloadbalancerresponse["status"] = Deleteloadbalancerresp.StatusCode
	Deleteloadbalancerresponse["body"] = string(body)
	resp = Deleteloadbalancerresponse
	return resp, err
}

//Listloadbalancer lists google loadbalancer pool.
func (googleloadbalancer *Googleloadbalancer) Listloadbalancer(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/beta/projects/" + options["Project"] + "/regions/" + options["Region"] + "/targetPools/"

//  url := "https://www.googleapis.com/compute/beta/projects/" + Project + "/regions/" + Region + "/targetPools"

	client := googleauth.SignJWT()

	Listloadbalancerrequest, err := http.NewRequest("GET", url, nil)
	Listloadbalancerrequest.Header.Set("Content-Type", "application/json")

	Listloadbalancerresp, err := client.Do(Listloadbalancerrequest)

	defer Listloadbalancerresp.Body.Close()

	body, err := ioutil.ReadAll(Listloadbalancerresp.Body)

	Listloadbalancerresponse := make(map[string]interface{})
	Listloadbalancerresponse["status"] = Listloadbalancerresp.StatusCode
	Listloadbalancerresponse["body"] = string(body)
	resp = Listloadbalancerresponse
	return resp, err
}

//Attachnodewithloadbalancer attach new google compute instance to google loadbalancer pool.
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
	url := "https://www.googleapis.com/compute/beta/projects/" + Project + "/regions/" + Region + "/targetPools/" + TargetPool + "/addInstance"

	Attachnodewithloadbalancerjsonmap := make(map[string]interface{})

	if len(Instances) != 0 {
		instance := []interface{}{}
		for i := 0; i < (len(Instances)); i++ {
			val := map[string]string{}
			val["instance"] = Instances[i]
			instance = append(instance, val)
		}
		Attachnodewithloadbalancerjsonmap["instances"] = instance
	}

	Attachnodewithloadbalancerjson, _ := json.Marshal(Attachnodewithloadbalancerjsonmap)

	Attachnodewithloadbalancerjsonstring := string(Attachnodewithloadbalancerjson)

	var Attachnodewithloadbalancerstringbyte = []byte(Attachnodewithloadbalancerjsonstring)

	client := googleauth.SignJWT()

	Attachnodewithloadbalancerrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Attachnodewithloadbalancerstringbyte))

	Attachnodewithloadbalancerrequest.Header.Set("Content-Type", "application/json")

	Attachnodewithloadbalancerresp, err := client.Do(Attachnodewithloadbalancerrequest)

	defer Attachnodewithloadbalancerresp.Body.Close()

	body, err := ioutil.ReadAll(Attachnodewithloadbalancerresp.Body)

	Attachnodewithloadbalancerresponse := make(map[string]interface{})
	Attachnodewithloadbalancerresponse["status"] = Attachnodewithloadbalancerresp.StatusCode
	Attachnodewithloadbalancerresponse["body"] = string(body)
	resp = Attachnodewithloadbalancerresponse
	return resp, err
}

//Detachnodewithloadbalancer Detach google compute instance from google loadbalancer pool.
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
	url := "https://www.googleapis.com/compute/beta/projects/" + Project + "/regions/" + Region + "/targetPools/" + TargetPool + "/removeInstance"

	Detachnodewithloadbalancerjsonmap := make(map[string]interface{})

	if len(Instances) != 0 {
		instance := []interface{}{}
		for i := 0; i < (len(Instances)); i++ {
			val := map[string]string{}
			val["instance"] = Instances[i]
			instance = append(instance, val)
		}
		Detachnodewithloadbalancerjsonmap["instances"] = instance
	}

	Detachnodewithloadbalancerjson, _ := json.Marshal(Detachnodewithloadbalancerjsonmap)

	Detachnodewithloadbalancerjsonstring := string(Detachnodewithloadbalancerjson)

	var Detachnodewithloadbalancerstringbyte = []byte(Detachnodewithloadbalancerjsonstring)

	client := googleauth.SignJWT()

	Detachnodewithloadbalancerrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Detachnodewithloadbalancerstringbyte))

	Detachnodewithloadbalancerrequest.Header.Set("Content-Type", "application/json")

	Detachnodewithloadbalancerresp, err := client.Do(Detachnodewithloadbalancerrequest)

	defer Detachnodewithloadbalancerresp.Body.Close()

	body, err := ioutil.ReadAll(Detachnodewithloadbalancerresp.Body)

	Detachnodewithloadbalancerresponse := make(map[string]interface{})
	Detachnodewithloadbalancerresponse["status"] = Detachnodewithloadbalancerresp.StatusCode
	Detachnodewithloadbalancerresponse["body"] = string(body)
	resp = Detachnodewithloadbalancerresponse
	return resp, err
}
