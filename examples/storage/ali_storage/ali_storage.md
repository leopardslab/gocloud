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

```
  createDisk := map[string]interface{}{
        "RegionId":    "cn-qingdao",
        "ZoneId":      "cn-qingdao-b",
        "Size":        20,
        "DiskName":    "ThisIsDiskName",
        "Description": "ThisIsDescription",
  }

  resp, err := alicloud.CreateDisk(createDisk)
  createDiskResp, err := alistorage.ParseCreateDiskResp(resp)
  fmt.Println(createDiskResp.StatusCode)
  fmt.Println(createDiskResp.DiskId)
```

or

```
    create, err := alistorage.NewCreateDiskBuilder().
        RegionID("cn-qingdao").
        ZoneID("cn-qingdao-b").
        Size(20).
        DiskName("ThisIsDiskName").
        Description("ThisIsDescription").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.CreateDisk(create)
    if err != nil {
        fmt.Println(err)
        return
    }
    createDiskResp, err := alistorage.ParseCreateDiskResp(resp)
    fmt.Println(createDiskResp.StatusCode)
    fmt.Println(createDiskResp.DiskId)
```

### Delete disk

```js
  deleteDisk := map[string]interface{}{
        "DiskId": "d-m5e2g66ws9m7r00ux87h",
  }

  resp, err := alicloud.DeleteDisk(deleteDisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    deleteDisk, err := alistorage.NewDeleteDiskBuilder().
        DiskID("d-m5e9fz4ninlgo4cdp85e").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.DeleteDisk(deleteDisk)
    if err != nil {
        fmt.Println(err)
        return
    }
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

  resp, err := alicloud.AttachDisk(attachDisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    attachDisk, err := alistorage.NewAttachDiskBuilder().
        InstanceID("i-m5e9rt5ix8dm8x96r8gw").
        DiskID("d-m5edrpqm9r8yzquyuzsg").
        DeleteWithInstance(false).
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.AttachDisk(attachDisk)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Detach disk

```js
  detachDisk := map[string]interface{}{
        "InstanceId": "i-m5e1lploaf58bddf0gah",
        "DiskId":     "d-m5edwiwlyyn7bwz6cdd4",
  }

  resp, err := alicloud.DetachDisk(detachDisk)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    detachDisk, err := alistorage.NewDetachDiskBuilder().
        InstanceID("i-m5e9rt5ix8dm8x96r8gw").
        DiskID("d-m5edrpqm9r8yzquyuzsg").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.DetachDisk(detachDisk)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Create snapshot

```js
  createsnapshot := map[string]interface{}{
        "DiskId":       "d-m5e7k6ycnx8b5zzsm0yp",
        "SnapshotName": "ThisIsSnapshotName",
  }

  resp, err := alicloud.CreateSnapshot(createsnapshot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    createSnapshot, err := alistorage.NewCreateSnapshotBuilder().
        DiskID("d-m5e3rdnsbhtjkr9idjsu").
        SnapshotName("ThisIsSnapshotName").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.CreateSnapshot(createSnapshot)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Delete snapshot

```js
  deletesnapshot := map[string]interface{}{
        "SnapshotId": "s-m5eave3s6oufpctcxynu",
  }

  resp, err := alicloud.DeleteSnapshot(deletesnapshot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    deleteSnapshot, err := alistorage.NewDeleteSnapshotBuilder().
        SnapshotID("s-m5ec7pkk6xpiff8o5hwl").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.DeleteSnapshot(deleteSnapshot)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```