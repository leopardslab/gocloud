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
    
    createBareMetalResp, err := vultrbaremetal.ParseCreateBareMetalResp(resp)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(createBareMetalResp.SUBID)
```

or

```
    createBareMetal, err := vultrbaremetal.NewCreateBareMetalBuilder().
        DCID(1).
        METALPLANID(100).
        OSID(127).
        Label("gocloud").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := vultr.BareMetal().CreateBareMetal(create)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
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

or

```
    deleteBareMetal, err := vultrbaremetal.NewDeleteBareMetalBuilder().
        SUBID(16917076).
        Build()
    if err != nil {
        fmt.Println(err)
        return
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

or

```
    haltBareMetal, err := vultrbaremetal.NewHaltBareMetalBuilder().
        SUBID(900000).
        Build()
    if err != nil {
        fmt.Println(err)
        return
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
    
    listBareMetalResp, err := vultrbaremetal.ParseListBareMetalResp(resp)
    if err != nil {
        fmt.Println(err)
    }
    for _, value := range listBareMetalResp.BareMetalSlice {
        fmt.Printf("%+v\n", value)
    }
```

or

```
    list, err := vultrbaremetal.NewListBareMetalBuilder().
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := vultr.BareMetal().ListBareMetal(list)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["status"])
    fmt.Println(response["body"])
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

or

```
    rebootBareMetal, err := vultrbaremetal.NewRebootBareMetalBuilder().
        SUBID(900000).
        Build()
    if err != nil {
        fmt.Println(err)
        return
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

or

```
    reinstallBareMetal, err := vultrbaremetal.NewReinstallBareMetalBuilder().
        SUBID(900000).
        Build()
    if err != nil {
        fmt.Println(err)
        return
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