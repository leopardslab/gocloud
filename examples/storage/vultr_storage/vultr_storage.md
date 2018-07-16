# gocloud storage - Vultr

## Configure Vultr credentials

Create `vultrconfig.json` as follows,
```
{
  "VultrAPIKey": "xxxxxxxxxxxx",
}
```

also You can setup environment variables as

```
export VultrAPIKey =  "xxxxxxxxxxxx"
```

## Initialize library

```
import "github.com/cloudlibz/gocloud/gocloud"

vultr, _ := gocloud.CloudProvider(gocloud.Vultrprovider)
```

### Create disk

```
    create := map[string]interface{}{
        "DCID":    1,
        "size_gb": 50,
        "label":   "test",
    }
    resp, err := vultr.CreateDisk(create)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    createDisk, err := vultrstorage.NewCreateDiskBuilder().
        DCID(1).
        SizeGb(50).
        Label("test").
        Build()
    resp, err := vultr.CreateDisk(create)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Delete disk

```
    deleteDisk := map[string]interface{}{
        "SUBID": 1313217,
    }
    resp, err := vultr.DeleteDisk(deleteDisk)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    deleteDisk, err := vultrstorage.NewDeleteDiskBuilder().
        SUBID(1313217).
        Build()
    resp, err := vultr.DeleteDisk(deleteDisk)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Attach disk

```
    attachDisk := map[string]interface{}{
        "SUBID":           1313217,
        "attach_to_SUBID": 1313207,
    }
    resp, err := vultr.AttachDisk(attachDisk)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    attachDisk, err := vultrstorage.NewAttachDiskBuilder().
        SUBID(1313217).
        AttachToSUBID(1313207).
        Build()
    resp, err := vultr.AttachDisk(attachDisk)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Detach disk

```
     detachDisk := map[string]interface{}{
         "SUBID": 1313217,
     }
     resp, err := vultr.DetachDisk(detachDisk)
     if err != nil {
         fmt.Println(err)
         return
     }
     response := resp.(map[string]interface{})
     fmt.Println(response["body"])
```

or

```
    detachDisk, err := vultrstorage.NewDetachDiskBuilder().
        SUBID(1313217).
        Build()
    resp, err := vultr.DetachDisk(detachDisk)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Create snapshot

```
    createSnapshot := map[string]interface{}{
        "SUBID": 1312965,
    }
    resp, err := vultr.CreateSnapshot(createSnapshot)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    createSnapshot, err := vultrstorage.NewCreateSnapshotBuilder().
        SUBID(1312965).
        Description("desc").
        Build()
    resp, err := vultr.CreateSnapshot(createSnapshot)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Delete snapshot

```
    deleteSnapshot := map[string]interface{}{
        "SNAPSHOTID": "5359435d28b9a",
    }
    resp, err := vultr.DeleteSnapshot(deleteSnapshot)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    deleteSnapshot, err := vultrstorage.NewDeleteSnapshotBuilder().
        SNAPSHOTID("5359435d28b9a").
        Build()
    resp, err := vultr.DeleteSnapshot(deleteSnapshot)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```