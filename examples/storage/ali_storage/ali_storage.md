# gocloud storage - Ali

## Configure Ali-cloud credentials

Create `alicloudconfig.json` as follows,
```js
{
  "AliAccessKeyID": "xxxxxxxxxxxx",
  "AliAccessKeySecret": "xxxxxxxxxxxx"
}
```

also You can setup environment variables as

```js
export AliAccessKeyID =  "xxxxxxxxxxxx"
export AliAccessKeySecret = "xxxxxxxxxxxx"
```

## Initialize library

```js
import "github.com/cloudlibz/gocloud/gocloud"

alicloud, _ := gocloud.CloudProvider(gocloud.Aliprovider)
```

### Create disk

```js
  createDisk := map[string]interface{}{
        "RegionId":    "cn-qingdao",
        "ZoneId":      "cn-qingdao-b",
        "Size":        20,
        "DiskName":    "ThisIsDiskName",
        "Description": "ThisIsDescription",
  }

  resp, err := alicloud.Createdisk(createDisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete disk

```js
  deleteDisk := map[string]interface{}{
        "DiskId": "d-m5e2g66ws9m7r00ux87h",
  }

  resp, err := alicloud.Deletedisk(deleteDisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Attach disk

```js
  attachDisk := map[string]interface{}{
        "InstanceId":         "i-m5e1lploaf58bddf0gah",
        "DiskId":             "d-m5edwiwlyyn7bwz6cdd4",
        "DeleteWithInstance": false,
  }

  resp, err := alicloud.Attachdisk(attachDisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Detach disk

```js
  detachDisk := map[string]interface{}{
		"InstanceId": "i-m5e1lploaf58bddf0gah",
        "DiskId":     "d-m5edwiwlyyn7bwz6cdd4",
  }

  resp, err := alicloud.Detachdisk(detachDisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Create snapshot

```js
  createsnapshot := map[string]interface{}{
		"DiskId":       "d-m5e7k6ycnx8b5zzsm0yp",
        "SnapshotName": "ThisIsSnapshotName",
  }

  resp, err := alicloud.Createsnapshot(createsnapshot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete snapshot

```js
  deletesnapshot := map[string]interface{}{
		"SnapshotId": "s-m5eave3s6oufpctcxynu",
  }

  resp, err := alicloud.Deletesnapshot(deletesnapshot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
