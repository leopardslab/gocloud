# gocloud DNS - Ali-cloud

## Configure Ali-cloud credentials

Create `alicloudconfig.json` as follows,
```
{
  "AliAccessKeyID": "xxxxxxxxxxxx",
  "AliAccessKeySecret": "xxxxxxxxxxxx"
}
```

also You can setup environment variables as

```
export AliAccessKeyID =  "xxxxxxxxxxxx"
export AliAccessKeySecret = "xxxxxxxxxxxx"
```

## Initialize library

```
import "github.com/cloudlibz/gocloud/gocloud"

alicloud, _ := gocloud.CloudProvider(gocloud.Aliprovider)
```

### Create DNS

```
  createDNS := map[string]interface{}{
          "DomainName": "oddcn.cn",
          "RR":         "test.gocloud",
          "Type":       "A",
          "Value":      "202.106.0.20",
          "TTL":        600,
          "Line":       "default",
  }

  resp, err := alicloud.CreateDns(createDNS)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete DNS

```js
  deleteDNS := map[string]interface{}{
          "RecordId": "3888946862348288",
  }

  resp, err := alicloud.DeleteDns(deleteDNS)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### List DNS

```
    list := map[string]interface{}{
        "DomainName": "oddcn.cn",
        "PageNumber": 1,
        "PageSize":   20,
    }
    resp, err := alicloud.ListDns(list)
    listDnsResp, err := alidns.ParseListDnsResp(resp)
    fmt.Printf("%+v", listDnsResp)
```
