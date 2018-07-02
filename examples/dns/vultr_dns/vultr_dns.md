# gocloud DNS - Vultr

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

### Create DNS

```
    createDNS := map[string]interface{}{
        "domain": "oddcn.cn",
        "name":   "gocloud.test1",
        "type":   "A",
        "data":   "192.0.2.1",
    }
    resp, err := vultr.CreateDns(createDNS)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Delete DNS

```
    deleteDNS := map[string]interface{}{
        "domain": "oddcn.cn",
        "RECORDID": 7065076,
    }
    resp, err := vultr.DeleteDns(deleteDNS)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### List DNS

```
    listDNS := map[string]interface{}{
        "domain": "oddcn.cn",
    }
    resp, err := vultr.ListDns(listDNS)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### List resource DNS record sets

```j
not support
```