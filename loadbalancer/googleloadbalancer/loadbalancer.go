package googleloadbalancer

import (
	"bytes"
	"encoding/json"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	//UnixDate reperesnts unix date-time format.
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	//RFC3339 reperesnts RFC3339 date-time format.
	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

//CreateLoadBalancer creates google loadbalancer pool.
func (googleloadbalancer *Googleloadbalancer) CreateLoadBalancer(request interface{}) (resp interface{}, err error) {

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

	CreateLoadBalancerjsonmap := make(map[string]interface{})

	CreateLoadBalancerdictnoaryconvert(option, CreateLoadBalancerjsonmap)

	CreateLoadBalancerjson, _ := json.Marshal(CreateLoadBalancerjsonmap)

	CreateLoadBalancerjsonstring := string(CreateLoadBalancerjson)

	var CreateLoadBalancerstringbyte = []byte(CreateLoadBalancerjsonstring)

	url := "https://www.googleapis.com/compute/beta/projects/" + Project + "/regions/" + Region + "/targetPools"

	client := googleauth.SignJWT()

	CreateLoadBalancerrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(CreateLoadBalancerstringbyte))

	CreateLoadBalancerrequest.Header.Set("Content-Type", "application/json")

	CreateLoadBalancerresp, err := client.Do(CreateLoadBalancerrequest)

	defer CreateLoadBalancerresp.Body.Close()

	body, err := ioutil.ReadAll(CreateLoadBalancerresp.Body)

	CreateLoadBalancerresponse := make(map[string]interface{})
	CreateLoadBalancerresponse["status"] = CreateLoadBalancerresp.StatusCode
	CreateLoadBalancerresponse["body"] = string(body)
	resp = CreateLoadBalancerresponse
	return resp, err
}

//DeleteLoadBalancer deletes google loadbalancer pool.
func (googleloadbalancer *Googleloadbalancer) DeleteLoadBalancer(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/beta/projects/" + options["Project"] + "/regions/" + options["Region"] + "/targetPools/" + options["TargetPool"]

	client := googleauth.SignJWT()

	DeleteLoadBalancerrequest, err := http.NewRequest("DELETE", url, nil)
	DeleteLoadBalancerrequest.Header.Set("Content-Type", "application/json")

	DeleteLoadBalancerresp, err := client.Do(DeleteLoadBalancerrequest)

	defer DeleteLoadBalancerresp.Body.Close()

	body, err := ioutil.ReadAll(DeleteLoadBalancerresp.Body)

	DeleteLoadBalancerresponse := make(map[string]interface{})
	DeleteLoadBalancerresponse["status"] = DeleteLoadBalancerresp.StatusCode
	DeleteLoadBalancerresponse["body"] = string(body)
	resp = DeleteLoadBalancerresponse
	return resp, err
}

//ListLoadBalancer lists google loadbalancer pool.
func (googleloadbalancer *Googleloadbalancer) ListLoadBalancer(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/beta/projects/" + options["Project"] + "/regions/" + options["Region"] + "/targetPools/"

	//  url := "https://www.googleapis.com/compute/beta/projects/" + Project + "/regions/" + Region + "/targetPools"

	client := googleauth.SignJWT()

	ListLoadBalancerrequest, err := http.NewRequest("GET", url, nil)
	ListLoadBalancerrequest.Header.Set("Content-Type", "application/json")

	ListLoadBalancerresp, err := client.Do(ListLoadBalancerrequest)

	defer ListLoadBalancerresp.Body.Close()

	body, err := ioutil.ReadAll(ListLoadBalancerresp.Body)

	ListLoadBalancerresponse := make(map[string]interface{})
	ListLoadBalancerresponse["status"] = ListLoadBalancerresp.StatusCode
	ListLoadBalancerresponse["body"] = string(body)
	resp = ListLoadBalancerresponse
	return resp, err
}

//AttachNodeWithLoadBalancer attach new google compute instance to google loadbalancer pool.
func (googleloadbalancer *Googleloadbalancer) AttachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error) {

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

	AttachNodeWithLoadBalancerjsonmap := make(map[string]interface{})

	if len(Instances) != 0 {
		instance := []interface{}{}
		for i := 0; i < (len(Instances)); i++ {
			val := map[string]string{}
			val["instance"] = Instances[i]
			instance = append(instance, val)
		}
		AttachNodeWithLoadBalancerjsonmap["instances"] = instance
	}

	AttachNodeWithLoadBalancerjson, _ := json.Marshal(AttachNodeWithLoadBalancerjsonmap)

	AttachNodeWithLoadBalancerjsonstring := string(AttachNodeWithLoadBalancerjson)

	var AttachNodeWithLoadBalancerstringbyte = []byte(AttachNodeWithLoadBalancerjsonstring)

	client := googleauth.SignJWT()

	AttachNodeWithLoadBalancerrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(AttachNodeWithLoadBalancerstringbyte))

	AttachNodeWithLoadBalancerrequest.Header.Set("Content-Type", "application/json")

	AttachNodeWithLoadBalancerresp, err := client.Do(AttachNodeWithLoadBalancerrequest)

	defer AttachNodeWithLoadBalancerresp.Body.Close()

	body, err := ioutil.ReadAll(AttachNodeWithLoadBalancerresp.Body)

	AttachNodeWithLoadBalancerresponse := make(map[string]interface{})
	AttachNodeWithLoadBalancerresponse["status"] = AttachNodeWithLoadBalancerresp.StatusCode
	AttachNodeWithLoadBalancerresponse["body"] = string(body)
	resp = AttachNodeWithLoadBalancerresponse
	return resp, err
}

//DetachNodeWithLoadBalancer Detach google compute instance from google loadbalancer pool.
func (googleloadbalancer *Googleloadbalancer) DetachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error) {
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

	DetachNodeWithLoadBalancerjsonmap := make(map[string]interface{})

	if len(Instances) != 0 {
		instance := []interface{}{}
		for i := 0; i < (len(Instances)); i++ {
			val := map[string]string{}
			val["instance"] = Instances[i]
			instance = append(instance, val)
		}
		DetachNodeWithLoadBalancerjsonmap["instances"] = instance
	}

	DetachNodeWithLoadBalancerjson, _ := json.Marshal(DetachNodeWithLoadBalancerjsonmap)

	DetachNodeWithLoadBalancerjsonstring := string(DetachNodeWithLoadBalancerjson)

	var DetachNodeWithLoadBalancerstringbyte = []byte(DetachNodeWithLoadBalancerjsonstring)

	client := googleauth.SignJWT()

	DetachNodeWithLoadBalancerrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(DetachNodeWithLoadBalancerstringbyte))

	DetachNodeWithLoadBalancerrequest.Header.Set("Content-Type", "application/json")

	DetachNodeWithLoadBalancerresp, err := client.Do(DetachNodeWithLoadBalancerrequest)

	defer DetachNodeWithLoadBalancerresp.Body.Close()

	body, err := ioutil.ReadAll(DetachNodeWithLoadBalancerresp.Body)

	DetachNodeWithLoadBalancerresponse := make(map[string]interface{})
	DetachNodeWithLoadBalancerresponse["status"] = DetachNodeWithLoadBalancerresp.StatusCode
	DetachNodeWithLoadBalancerresponse["body"] = string(body)
	resp = DetachNodeWithLoadBalancerresponse
	return resp, err
}
