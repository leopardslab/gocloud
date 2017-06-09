package gce

import(
  "fmt"
  //"net"
  "net/http"
  "io/ioutil"
  "bytes"
)




func (gce *GCE)describenode(options interface{})(resp interface{},err error){


  token := sign()

  client := &http.Client{}

  req, _ := http.NewRequest("GET", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances", nil)

  token.SetAuthHeader(req)

  res,_ := client.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
  return
}

func (gce *GCE)Createnode(options interface{})(resp interface{},err error){


	 token := sign()

  	client := &http.Client{}

	var jsonStr2 = []byte(`{
  "disks":[
    {
      "boot":true,
      "initializeParams":{
        "sourceImage":"https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301"
      }
    }
  ],
  "machineType":"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
  "name":"gocloud-100",
  "networkInterfaces":[
    {
      "accessConfigs":[
        {
          "name":"external-nat",
          "type":"ONE_TO_ONE_NAT"
        }
      ],
      "network":"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default"
    }
  ],
  "zone":"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c"
}`)

	req2, err := http.NewRequest("POST", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances", bytes.NewBuffer(jsonStr2))

	req2.Header.Set("Content-Type", "application/json")

	//req2.Header.Set("Content-length", "743")

	token.SetAuthHeader(req2)

	resp2, err := client.Do(req2)

	defer resp2.Body.Close()

	body2, err := ioutil.ReadAll(resp2.Body)

	fmt.Println(string(body2))


	return

}



func (gce *GCE)deletenode(options interface{})(resp interface{},err error){

req5, err := http.NewRequest("DELETE", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances/instance-10", nil)


	req5.Header.Set("Content-Type", "application/json")

	//req2.Header.Set("Content-length", "743")

	tok.SetAuthHeader(req5)

	resp5, err := client.Do(req5)

	defer resp4.Body.Close()

	body5, err := ioutil.ReadAll(resp5.Body)

	fmt.Println("\n\n\n\n\n\n\n\n\n\n")

	fmt.Println(string(body5))



}





func (gce *GCE)stopnode(options interface{})(resp interface{},err error){

	req4, err := http.NewRequest("POST", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances/instance-10/stop", nil)


	req3.Header.Set("Content-Type", "application/json")

	//req2.Header.Set("Content-length", "743")

	tok.SetAuthHeader(req4)

	resp4, err := client.Do(req4)

	defer resp4.Body.Close()

	body4, err := ioutil.ReadAll(resp4.Body)

	fmt.Println("\n\n\n\n\n\n\n\n\n\n")

	fmt.Println(string(body4))


}




func (gce *GCE)startnode(options interface{})(resp interface{},err error){

	req3, err := http.NewRequest("POST", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances/instance-11/start", nil)


	req3.Header.Set("Content-Type", "application/json")

	//req2.Header.Set("Content-length", "743")

	tok.SetAuthHeader(req3)

	resp3, err := client.Do(req3)

	defer resp3.Body.Close()

	body3, err := ioutil.ReadAll(resp3.Body)

	fmt.Println("\n\n\n\n\n\n\n\n\n\n")

	fmt.Println(string(body3))


	fmt.Println("\n\n\n\n\n\n\n\n\n\n")

}
