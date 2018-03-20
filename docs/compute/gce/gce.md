# gocloud compute - gce

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

### Create instance

```js
 InitializeParams := map[string]string{
    "SourceImage": "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
    "DiskType":    "projects/sheltermap-1493101612061/zones/us-east4-c/diskTypes/pd-standard",
    "DiskSizeGb":  "10",
  }

  disk := []map[string]interface{}{
     {
	"Boot":  true,
	"AutoDelete":   false,
	"DeviceName":   "DeviceName",
	"Type":         "PERSISTENT",
	"Mode":         "READ_WRITE",
	"InitializeParams": InitializeParams,
      },
   }

 AccessConfigs := []map[string]string{{
	"Name": "external-nat",
	"Type": "ONE_TO_ONE_NAT",
  },
  }

  NetworkInterfaces := []map[string]interface{}{
   {
	"Network": "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
	"Subnetwork":    "projects/sheltermap-1493101612061/regions/us-east4/subnetworks/default",
        "AccessConfigs": AccessConfigs,
    },
   }

   createnode := map[string]interface{}{
      "projectid":   "sheltermap-1493101612061",
	  "Name" :   "testing-scorelab",
    "MachineType":    "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
    "Zone":        "us-central1-b",
    "disk":        disk,
    "NetworkInterfaces": NetworkInterfaces,
    }

  resp, err := googlecloud.Createnode(createnode)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])

  ```

### Stop instance

```js
  stopnode := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "testing-scorelab2",
		"Zone":      "us-west1-c",
	}

 resp , err := googlecloud.Stopnode(stopnode)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Start instance

```js
  stop := map[string]string{
	   "projectid":"sheltermap-1493101612061",
	   "instance":"instance-10",
	    "Zone": "us-west1-c",

	   }

  resp,err := googlecloud.Stopnode(stop)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Reboot instance

```js
 	reboot := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "testing-scorelab",
		"Zone":      "us-west1-c",
	}

  resp, err := googlecloud.Rebootnode(reboot)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete instance

```js
 deletenode := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "testing-scorelab2",
		"Zone":      "us-west1-c",
	}

  resp, err := googlecloud.Deletenode(deletenode)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```