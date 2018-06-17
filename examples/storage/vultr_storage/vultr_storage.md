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
    resp, err := vultr.Createdisk(create)
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
      resp, err := vultr.Deletedisk(deleteDisk)
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
        "SUBID":           1313217,
        "attach_to_SUBID": 1313207,
    }
    resp, err := vultr.Attachdisk(attachDisk)
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
     resp, err := vultr.Detachdisk(detachDisk)
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
      resp, err := vultr.Createsnapshot(createSnapshot)
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
      resp, err := vultr.Deletesnapshot(deleteSnapshot)
      if err != nil {
          fmt.Println(err)
          return
      }
      response := resp.(map[string]interface{})
      fmt.Println(response["body"])
```
