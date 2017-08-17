# gocloud storage - AWS

## Configure AWS credentials

Create `gocloudconfig.json` as follows,
```js
{
  "AWSAccessKeyID": "xxxxxxxxxxxx",
  "AWSSecretKey": "xxxxxxxxxxxx",
}
```

also You can setup enviroment variables as

```js
export AWSAccessKeyID =  "xxxxxxxxxxxx"
export AWSSecretKey = "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/scorelab/gocloud-v2/gocloud"

amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
```

### Create disk

```js
  createdisk := map[string]interface{}{
		"AvailZone":  "us-east-1a",
		"VolumeSize": 100,
		"Region":     "us-east-1",
	}
  
  resp, err := amazonstorage.Createdisk(createdisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

###  delete disk

```js
  deletedisk := map[string]string{
		"VolumeId": "vol-0996a16ff8f032760",
		"Region":   "us-east-1",
	}
 
  resp, err := amazonstorage.Deletedisk(deletedisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### attach disk

```js
  attachdisk := map[string]string{
		"VolumeId":   "vol-049426a70587418d7",
		"InstanceId": "i-0050d952f9b8d45d5",
		"Device":     "/dev/sdh",
		"Region":     "us-east-1",
	}
	
  resp, err := amazonstorage.Attachdisk(attachdisk)
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

  resp, err := amazonstorage.Detachdisk(detachdisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```


### create snapshot

```js
createsnapshot := map[string]string{
		"VolumeId":    "vol-047d011f7536d2b7c",
		"Description": "create snapshot for vol-047d011f7536d2b7c",
		"Region":      "us-east-1",
	}
	
  resp, err := amazonstorage.Createsnapshot(createsnapshot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### delete snapshot

```js
  deletesnapshot := map[string]string{
		"SnapshotId": "snap-0f0839076356ce6cb",
		"Region":     "us-east-1",
	}
  
  resp, err := amazonstorage.Deletesnapshot(deletesnapshot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
