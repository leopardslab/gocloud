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

  resp, err := alicloud.Createdns(createDNS)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete DNS

```js
  deleteDNS := map[string]interface{}{
  		"RecordId": "3888946862348288",
  }

  resp, err := alicloud.Deletedns(deleteDNS)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### List DNS

```js
  listDNS := map[string]interface{}{
  		"PageNumber": 1,
  		"PageSize":   20,
  }

  resp, err := alicloud.Listdns(listDNS)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### List resource DNS record sets

```j
  listResourceDNSRecordSets := map[string]interface{}{
  		"DomainName": "oddcn.cn",
  }

  resp, err := alicloud.ListResourcednsRecordSets(listResourceDNSRecordSets)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```