# gocloud compute - Vultr server

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

### Create instance

```
    create := map[string]interface{}{
        "DCID":      1,
        "VPSPLANID": 201,
        "OSID":      127,
    }
    resp, err := vultrCloud.Createnode(create)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Start instance

```
    start := map[string]interface{}{
        "SUBID":      16183500,
    }
    resp, err := vultrCloud.Startnode(start)
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
```

### Stop instance

```
// This API is not provided by Vultr 
```

### Reboot instance

```
    reboot := map[string]interface{}{
        "SUBID": 16183500,
    }
    resp, err := vultrCloud.Rebootnode(reboot)
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
```

### Delete instance

```
    d := map[string]interface{}{
        "SUBID": 16183500,
    }
    resp, err := vultrCloud.Deletenode(d)
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
```