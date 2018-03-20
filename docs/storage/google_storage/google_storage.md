# gocloud storage - gce

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

### Create disk

```js
 createdisk := map[string]interface{}{
		"projectid": "sheltermap-1493101612061",
		"Name":      "disk-11",
		"Type":      "pd-standard",
		"Zone":      "us-east4-c",
	}

  resp, err := googlecloud.Createdisk(createdisk)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])

  ```

### Attach disk

```js
 attachdisk := map[string]string{
    "projectid":  "sheltermap-1493101612061",
    "instance":   "testing-scorelab",
    "Zone":       "us-east4-c",
    "deviceName": "projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
 }

 resp, err := googlecloud.Attachdisk(attachdisk)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
### Delete disk

```js
  deletedisk := map[string]string{
		"VolumeId": "vol-05371175f10d2e6a4",
	}

 resp, err := googlecloud.Deletedisk(deletedisk)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Detach disk

```js
 detachdisk := map[string]string{
	"projectid":  "sheltermap-1493101612061",
	"instance":   "testing-scorelab",
	"Zone":       "us-east4-c",
	"deviceName": "projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
 }

 resp, err := googlecloud.Detachdisk(detachdisk)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Create snapshot

```js
 createsnapshot := map[string]interface{}{
		"projectid": "sheltermap-1493101612061",
		"disk":      "disk-11",
		"Zone":      "us-east4-c",
		"Name":      "disk-12",
	}

  resp, err := googlecloud.Createsnapshot(createsnapshot)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete snapshot

```js
 deletesnapshot := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"snapshot":  "disk-12",
	}

  resp, err := googlecloud.Deletesnapshot(deletesnapshot)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```