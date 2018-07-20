# gocloud bare metal - Vultr

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

vultr := gocloud.VultrProvider()
```

### Create

```
    create := map[string]interface{}{
        "DCID":        1,
        "METALPLANID": 100,
        "OSID":        127,
    }
    resp, err := vultr.BareMetal().CreateBareMetal(create)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
    
    createBareMetalResp, err := vultrbaremetal.ParseCreateBareMetalResp(response["body"])
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(createBareMetalResp.SUBID)
```

### Delete

```
    deleteBareMetal := map[string]interface{}{
        "SUBID": 900000,
    }
    resp, err := vultr.BareMetal().DeleteBareMetal(deleteBareMetal)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
```

### Halt

```
    haltBareMetal := map[string]interface{}{
        "SUBID": 900000,
    }
    resp, err := vultr.BareMetal().HaltBareMetal(haltBareMetal)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
```

### List

```
    resp, err := vultr.BareMetal().ListBareMetal(nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
    
    
    listBareMetalResp, err := vultrbaremetal.ParseListBareMetalResp(response["body"])
    if err != nil {
        fmt.Println(err)
    }
    for _, bareMetal := range listBareMetalResp {
        fmt.Printf("%+v", bareMetal)
    }
```

### Reboot

```
    rebootBareMetal := map[string]interface{}{
        "SUBID": 900000,
    }
    resp, err := vultr.BareMetal().RebootBareMetal(rebootBareMetal)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
```

### Reinstall

```
    reinstallBareMetal := map[string]interface{}{
        "SUBID": 900000,
    }
    resp, err := vultr.BareMetal().ReinstallBareMetal(reinstallBareMetal)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
```