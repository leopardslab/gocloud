package digioceanloadbalancer

import (
  digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
  "bytes"
  "fmt"
  "io/ioutil"
	"net/http"
  "encoding/json"
  "errors"
  "strconv"
)

// loadBalancerBasePath is the endpoint URL for digitalocean API.
const loadBalancerBasePath = "https://api.digitalocean.com/v2/load_balancers"

// Creatloadbalancer function creates a new load balancer.
func (digioceanloadbalancer *LoadBalancer) Creatloadbalancer(request interface{}) (resp interface{}, err error) {

	var loadBalancerInstance LoadBalancer	// Initialize LoadBalancer struct
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	param := make(map[string]interface{})
  param = request.(map[string]interface{})

  for key, value := range param {

    switch key {

      case "Name":
        name, _ := value.(string)
        loadBalancerInstance.Name = name

      case "Algorithm":
        algorithm, _ := value.(string)
        loadBalancerInstance.Algorithm = algorithm

      case "Region":
        region, _ := value.(string)
        loadBalancerInstance.Region = region

      case "ForwardingRules":
        forwardingRulesParam, _ := value.([]map[string]interface{})
        for i := 0; i < len(forwardingRulesParam); i++ {
          var loadBalancerForwardingRule ForwardingRule
          for forwardingRulesParamKey, forwardingRulesParamValue := range forwardingRulesParam[i] {
            switch forwardingRulesParamKey {
              case "EntryProtocol":
                loadBalancerForwardingRule.EntryProtocol = forwardingRulesParamValue.(string)
              case "EntryPort":
                loadBalancerForwardingRule.EntryPort = forwardingRulesParamValue.(int)
              case "TargetProtocol":
                loadBalancerForwardingRule.TargetProtocol = forwardingRulesParamValue.(string)
              case "TargetPort":
                loadBalancerForwardingRule.TargetPort = forwardingRulesParamValue.(int)
              case "CertificateID":
                loadBalancerForwardingRule.CertificateID = forwardingRulesParamValue.(string)
              case "TlsPassthrough":
                loadBalancerForwardingRule.TlsPassthrough = forwardingRulesParamValue.(bool)
            }
          }
          loadBalancerInstance.ForwardingRules = append(loadBalancerInstance.ForwardingRules, loadBalancerForwardingRule)
        }

      case "HealthCheck":
      healthCheckParam, _ := value.(map[string]interface{})
      for key, value := range healthCheckParam {
        switch key {
          case "Protocol":
            protocol, _ := value.(string)
            loadBalancerInstance.HealthCheck.Protocol = protocol
          case "Port":
            port, _ := value.(int)
            loadBalancerInstance.HealthCheck.Port = port
          case "Path":
            path, _ := value.(string)
            loadBalancerInstance.HealthCheck.Path = path
          case "CheckIntervalSeconds":
            checkIntervalSeconds, _ := value.(int)
            loadBalancerInstance.HealthCheck.CheckIntervalSeconds = checkIntervalSeconds
          case "ResponseTimeoutSeconds":
            responseTimeoutSeconds, _ := value.(int)
            loadBalancerInstance.HealthCheck.ResponseTimeoutSeconds = responseTimeoutSeconds
          case "HealthyThreshold":
            healthyThreshold, _ := value.(int)
            loadBalancerInstance.HealthCheck.HealthyThreshold = healthyThreshold
          case "UnhealthyThreshold":
            unhealthyThreshold, _ := value.(int)
            loadBalancerInstance.HealthCheck.UnhealthyThreshold = unhealthyThreshold
          }
        }

      case "StickySessions":
      stickySessionsParam, _ := value.(map[string]interface{})
      for key, value := range stickySessionsParam {
        switch key {
          case "Type":
            typeStickySessions, _ := value.(string)
            loadBalancerInstance.StickySessions.Type = typeStickySessions
          case "CookieName":
            cookieName, _ := value.(string)
            loadBalancerInstance.StickySessions.CookieName = cookieName
          case "CookieTtlSeconds":
            cookieTTLSeconds, _ := value.(int)
            loadBalancerInstance.StickySessions.CookieTtlSeconds = cookieTTLSeconds
          }
        }

      case "DropletIDs":
        dropletIds, _ := value.([]int)
        loadBalancerInstance.DropletIDs = dropletIds

      case "Tag":
        tag, _ := value.(string)
        loadBalancerInstance.Tag = tag

      case "RedirectHttpToHttps":
        redirectHTTPToHTTPS, _ := value.(bool)
        loadBalancerInstance.RedirectHttpToHttps = redirectHTTPToHTTPS

    } // Closes switch-case

  } // Closes for loop

  loadBalancerInstanceJSON, _ := json.Marshal(loadBalancerInstance)
  loadBalancerInstanceJSONString := string(loadBalancerInstanceJSON)
  var loadBalancerInstanceJSONStringbyte = []byte(loadBalancerInstanceJSONString)

  createLoadBalancerReq, err := http.NewRequest("POST", loadBalancerBasePath, bytes.NewBuffer(loadBalancerInstanceJSONStringbyte))
  if err != nil {
    fmt.Println(err)
  }
  createLoadBalancerReq.Header.Set("Content-Type", "application/json")
  createLoadBalancerReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	createLoadBalancerResp, err := http.DefaultClient.Do(createLoadBalancerReq)
	if err != nil {
		fmt.Println(err)
	}

	defer createLoadBalancerResp.Body.Close()

	responseBody, err := ioutil.ReadAll(createLoadBalancerResp.Body)
  createLoadBalancerResponse := make(map[string]interface{})
	createLoadBalancerResponse["status"] = createLoadBalancerResp.StatusCode
	createLoadBalancerResponse["body"] = string(responseBody)
	resp = createLoadBalancerResponse

	return resp, err

}

// Deleteloadbalancer function deletes a load balancer.
func (digioceanloadbalancer *LoadBalancer) Deleteloadbalancer(request interface{}) (resp interface{}, err error) {

  options := request.(map[string]string)
  inputID, err := strconv.Atoi(options["ID"])
  if err != nil {
    fmt.Println(err)
  }
  if inputID < 1 {
		return nil, errors.New("loadBalancerID cannot be less than 1")
	}

	url := loadBalancerBasePath + "/" + options["ID"]
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	deleteLoadBalancerReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	deleteLoadBalancerReq.Header.Set("Content-Type", "application/json")
	deleteLoadBalancerReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	deleteLoadBalancerResp, err := http.DefaultClient.Do(deleteLoadBalancerReq)
	if err != nil {
		fmt.Println(err)
	}

	defer deleteLoadBalancerResp.Body.Close()

	responseBody, err := ioutil.ReadAll(deleteLoadBalancerResp.Body)
	deleteLoadBalancerResponse := make(map[string]interface{})
	deleteLoadBalancerResponse["status"] = deleteLoadBalancerResp.StatusCode
	deleteLoadBalancerResponse["body"] = string(responseBody)
	resp = deleteLoadBalancerResponse

	return resp, err

}

// Listloadbalancer function lists load balancers.
func (digioceanloadbalancer *LoadBalancer) Listloadbalancer(request interface{}) (resp interface{}, err error) {

  DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

  listLoadBalancerReq, err := http.NewRequest("GET", loadBalancerBasePath, nil)
  if err != nil {
    fmt.Println(err)
  }
  listLoadBalancerReq.Header.Set("Content-Type", "application/json")
  listLoadBalancerReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

  listLoadBalancerResp, err := http.DefaultClient.Do(listLoadBalancerReq)
  if err != nil {
    fmt.Println(err)
  }

  defer listLoadBalancerResp.Body.Close()

  responseBody, err := ioutil.ReadAll(listLoadBalancerResp.Body)
  listLoadBalancerResponse := make(map[string]interface{})
  listLoadBalancerResponse["status"] = listLoadBalancerResp.StatusCode
  listLoadBalancerResponse["body"] = string(responseBody)
  resp = listLoadBalancerResponse

  return resp, err

}

// Attachnodewithloadbalancer function attaches a load balancer to a droplet.
func (digioceanloadbalancer *LoadBalancer) Attachnodewithloadbalancer(request interface{}) (resp interface{}, err error) {

  DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

  var loadBalancerID int  // To store the loadBalancerID
  var DropletIDs []int  // To store the DropletIDs

  param := make(map[string]interface{})
  param = request.(map[string]interface{})

  for key, value := range param {

    switch key {
      case "LoadBalancerID":
        loadBalancerIDvalue, _ := value.(int)
        loadBalancerID = loadBalancerIDvalue

      case "DropletIDs":
        dropletIds, _ := value.([]int)
        DropletIDs = dropletIds
    } // Closes switch-case

  } // Closes for loop

  url := loadBalancerBasePath + "/" + loadBalancerID + "/droplets"

  loadBalancerInstance := make(map[string]interface{})

	if len(DropletIDs) != 0 {
		instance := []interface{}{}
		for i := 0; i < (len(DropletIDs)); i++ {
			val := map[string]string{}
			val["droplet_ids"] = DropletIDs[i]
			instance = append(instance, val)
		}
		loadBalancerInstance["droplet_ids"] = instance
	}

  loadBalancerInstanceJSON, _ := json.Marshal(loadBalancerInstance)
  loadBalancerInstanceJSONString := string(loadBalancerInstanceJSON)
  var loadBalancerInstanceJSONStringbyte = []byte(loadBalancerInstanceJSONString)

  attachDropeltToLoadBalancerReq, err := http.NewRequest("POST", url, bytes.NewBuffer(loadBalancerInstanceJSONStringbyte))
  if err != nil {
    fmt.Println(err)
  }
  attachDropeltToLoadBalancerReq.Header.Set("Content-Type", "application/json")
  attachDropeltToLoadBalancerReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

  attachDropeltToLoadBalancerResp, err := http.DefaultClient.Do(attachDropeltToLoadBalancerReq)
  if err != nil {
    fmt.Println(err)
  }

  defer attachDropeltToLoadBalancerResp.Body.Close()

  responseBody, err := ioutil.ReadAll(attachDropeltToLoadBalancerResp.Body)
  attachDropeltToLoadBalancerResponse := make(map[string]interface{})
  attachDropeltToLoadBalancerResponse["status"] = attachDropeltToLoadBalancerResp.StatusCode
  attachDropeltToLoadBalancerResponse["body"] = string(responseBody)
  resp = attachDropeltToLoadBalancerResponse

  return resp, err
}

// Detachnodewithloadbalancer function detaches a load balancer from a droplet.
func (digioceanloadbalancer *LoadBalancer) Detachnodewithloadbalancer(request interface{}) (resp interface{}, err error) {

  DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

  var loadBalancerID int  // To store the loadBalancerID
  var DropletIDs []int  // To store the DropletIDs

  param := make(map[string]interface{})
  param = request.(map[string]interface{})

  for key, value := range param {

    switch key {
      case "LoadBalancerID":
        loadBalancerIDvalue, _ := value.(int)
        loadBalancerID = loadBalancerIDvalue

      case "DropletIDs":
        dropletIds, _ := value.([]int)
        DropletIDs = dropletIds
    } // Closes switch-case

  } // Closes for loop

  url := loadBalancerBasePath + "/" + loadBalancerID + "/droplets"

  loadBalancerInstance := make(map[string]interface{})

  if len(DropletIDs) != 0 {
    instance := []interface{}{}
    for i := 0; i < (len(DropletIDs)); i++ {
      val := map[string]string{}
      val["droplet_ids"] = DropletIDs[i]
      instance = append(instance, val)
    }
    loadBalancerInstance["droplet_ids"] = instance
  }

  loadBalancerInstanceJSON, _ := json.Marshal(loadBalancerInstance)
  loadBalancerInstanceJSONString := string(loadBalancerInstanceJSON)
  var loadBalancerInstanceJSONStringbyte = []byte(loadBalancerInstanceJSONString)

  detachDropeltToLoadBalancerReq, err := http.NewRequest("DELETE", url, bytes.NewBuffer(loadBalancerInstanceJSONStringbyte))
  if err != nil {
    fmt.Println(err)
  }
  detachDropeltToLoadBalancerReq.Header.Set("Content-Type", "application/json")
  detachDropeltToLoadBalancerReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

  detachDropeltToLoadBalancerResp, err := http.DefaultClient.Do(detachDropeltToLoadBalancerReq)
  if err != nil {
    fmt.Println(err)
  }

  defer detachDropeltToLoadBalancerResp.Body.Close()

  responseBody, err := ioutil.ReadAll(detachDropeltToLoadBalancerResp.Body)
  detachDropeltToLoadBalancerResponse := make(map[string]interface{})
  detachDropeltToLoadBalancerResponse["status"] = detachDropeltToLoadBalancerResp.StatusCode
  detachDropeltToLoadBalancerResponse["body"] = string(responseBody)
  resp = detachDropeltToLoadBalancerResponse

  return resp, err
}
