# gocloud container - gce

## Configure Google Cloud credentials.

Download your service account credentials file from Google Cloud and save it as `googlecloudinfo.json` in your <b>HOME</b> directory.

You can also set the credentials as environment variables:
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

import "github.com/shlokgilda/gocloud/gocloud"

googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
```

### Create cluster

```js
  nodepools := []map[string]interface{}{
       {
	"name":  "default-pool",
	"initialNodeCount": 3,
	},
  }

  createcluster := map[string]interface{}{
    "Project":   "sheltermap-1493101612061",
     "Name"  :   "cluster-2",
     "Zone"  :    "us-central1-a",
    "nodePools" : nodepools,
  }

  resp, err := googlecloud.Createcluster(createcluster)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])

  ```

### Delete cluster

```js
 deletecluster := map[string]string{
		"Project":   "sheltermap-1493101612061",
		"clusterId": "cluster-1",
		"Zone":      "us-central1-a",
	}

 resp , err := googlecloud.Deletecluster(deletecluster)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
### Create service

```js

  createservice := map[string]interface{}{
	"Project"   :   "sheltermap-1493101612061",
	"clusterId" :   "cluster-2",
	"Zone"      :   "us-central1-a",
	"Name"      :   "nodepool",
	}

  resp, err := googlecloud.Createservice(createservice)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete service

```js
deleteservice := map[string]string{
		"Project":    "sheltermap-1493101612061",
		"clusterId":  "cluster-2",
		"Zone":       "us-central1-a",
		"nodePoolId": "nodepool",
	}

 resp, err := googlecloud.Deleteservice(deleteservice)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```