# gocloud loadbancer - DigitalOcean

## Configure DigitalOcean credentials.

Create `digioceancloudconfig.json` in your <b>HOME/.gocloud</b> directory as follows:
```js
{
  "DigiOceanAccessToken": "xxxxxxxxxxxx"
}
```

You can also set the credentials as environment variables:
```js
export DigiOceanAccessToken =  "xxxxxxxxxxxx"
```


## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

digioceancloud, _ := gocloud.CloudProvider(gocloud.Digioceanprovider)
```

### Create disk

```js
  create := map[string]interface{}{
    "Name": "example-01",
    "Region": "nyc3",
    "Description":  "Block store for examples",
    "SizeGigaBytes":  1,
    "SnapshotID": nil,
  }

 resp, err := digioceancloud.CreateDisk(create)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete disk

```js
  delete1 := map[string]string{
    "VolumeID": "7724db7c-e098-11e5-b522-000f53304e51",
   }

  resp, err := digioceancloud.DeleteDisk(delete1)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Create snapshot

```js
  create := map[string]interface{}{
    "VolumeID": "7724db7c-e098-11e5-b522-000f53304e5",
    "SnapshotName": "big-data-snapshot1475261774",
  }

  resp, err := digioceancloud.CreateSnapshot(create)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete snapshot

```js
  delete1 := map[string]string{
    "SnapshotID": "7724db7c-e098-11e5-b522-000f53304e51",
   }

  resp, err := digioceancloud.DeleteSnapshot(delete1)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Attach disk to droplet

```js
  create := map[string]interface{}{
    "VolumeID": "7724db7c-e098-11e5-b522-000f53304e51",
    "DropletID": 9978454,
    "Region":  "nyc3",
  }

  resp, err := digioceancloud.AttachDisk(create)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Detach disk from droplet

```js
  delete1 := map[string]interface{}{
    "VolumeID": "7724db7c-e098-11e5-b522-000f53304e51",
    "DropletID": 9978454,
    "Region":  "nyc3",
  }

 resp, err := digioceancloud.DetachDisk(delete1)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
