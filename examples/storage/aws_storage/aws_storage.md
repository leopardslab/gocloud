# gocloud storage - AWS

## Configure AWS credentials.

Create `amazoncloudconfig.json` in your <b>HOME/.gocloud</b> directory as follows:
```js
{
  "AWSAccessKeyID": "xxxxxxxxxxxx",
  "AWSSecretKey": "xxxxxxxxxxxx"
}
```

You can also set the credentials as environment variables:
```js
export AWSAccessKeyID =  "xxxxxxxxxxxx"
export AWSSecretKey = "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
```

### Create disk

```js
  createdisk := map[string]interface{}{
		"AvailZone":  "us-east-1a",
		"VolumeSize": 100,
		"Region":     "us-east-1",
	}

  resp, err := amazonstorage.CreateDisk(createdisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete disk

```js
  deletedisk := map[string]string{
		"VolumeId": "vol-0996a16ff8f032760",
		"Region":   "us-east-1",
	}

  resp, err := amazonstorage.DeleteDisk(deletedisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Attach disk

```js
  attachdisk := map[string]string{
		"VolumeId":   "vol-049426a70587418d7",
		"InstanceId": "i-0050d952f9b8d45d5",
		"Device":     "/dev/sdh",
		"Region":     "us-east-1",
	}

  resp, err := amazonstorage.AttachDisk(attachdisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Detach disk

```js
 detachdisk := map[string]string{
		"VolumeId":   "vol-049426a70587418d7",
		"InstanceId": "i-0050d952f9b8d45d5",
		"Device":     "/dev/sdh",
		"Force":      "true",
		"Region":     "us-east-1",
	}

  resp, err := amazonstorage.DetachDisk(detachdisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```


### Create snapshot

```js
createsnapshot := map[string]string{
		"VolumeId":    "vol-047d011f7536d2b7c",
		"Description": "create snapshot for vol-047d011f7536d2b7c",
		"Region":      "us-east-1",
	}

  resp, err := amazonstorage.CreateSnapshot(createsnapshot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete snapshot

```js
  deletesnapshot := map[string]string{
		"SnapshotId": "snap-0f0839076356ce6cb",
		"Region":     "us-east-1",
	}

  resp, err := amazonstorage.DeleteSnapshot(deletesnapshot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
