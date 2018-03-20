# gocloud loadbalancer - gce

## Configure google credentials


download service account credential file from google cloud save as `googlecloudinfo.json`,


also You can setup enviroment variables as

```js
export PrivateKey =  "xxxxxxxxxxxx"
export Type =  "xxxxxxxxxxxx"
export ProjectID = "xxxxxxxxxxxx"
export PrivateKeyID = "xxxxxxxxxxxx"
export ClientEmail = "xxxxxxxxxxxx"
export ClientID = "xxxxxxxxxxxx"
export AuthURI = "xxxxxxxxxxxx"
export TokenURI = "xxxxxxxxxxxx"
export AuthProviderX509CertURL = "xxxxxxxxxxxx"
export ClientX509CertURL =  "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
```

### Create loadbalancer

```js
 creatloadbalancer := map[string]interface{}{
   "Project":  "sheltermap-1493101612061",
   "Name"   :  "google-loadbalancer",
   "Region":   "us-central1",
   "Instances": []string{"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-central1-b/instances/instance-1"},
	}

  resp, err := googlecloud.Creatloadbalancer(creatloadbalancer)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])

  ```

### List loadbalancer

```js
 listloadbalancer := map[string]string{
     "Project": "sheltermap-1493101612061",
     "Region":  "us-central1",
 }

 resp, err := googlecloud.Listloadbalancer(listloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
### attach node with loadbalancer

```js

  attachnodewithloadbalancer := map[string]interface{}{
	"Project":    "sheltermap-1493101612061",
	"Region":     "us-central1",
	"TargetPool": "google-loadbalancer",
	"Instances":  []string{"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-central1-b/instances/instance-2"},
	}

 resp, err := googlecloud.Attachnodewithloadbalancer(attachnodewithloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### detach node with loadbalancer

```js
  detachnodewithloadbalancer := map[string]interface{}{
      "Project":    "sheltermap-1493101612061",
      "Region":     "us-central1",
      "TargetPool": "google-loadbalancer",
      "Instances":  []string{"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-central1-b/instances/instance-2"},
	}

  resp, err := googlecloud.Detachnodewithloadbalancer(detachnodewithloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```


### delete loadbalancer
```js
  deleteloadbalancer := map[string]string{
		"Project":    "sheltermap-1493101612061",
		"Region":     "us-central1",
		"TargetPool": "google-loadbalancer",
	}

 resp, err := googlecloud.Deleteloadbalancer(deleteloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
