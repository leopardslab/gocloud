# gocloud compute - DigitalOcean

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

### Create DNS record

```js
  create := map[string]interface{}{
    "DomainName": "example.com",
    "Type":       "A",
    "Name":       "www",
    "Data":       "162.10.66.0",
    "Priority":   nil,
    "Port":       nil,
    "TimeToLive": 1800,
    "Weight":     nil,
    "Flags":      nil,
    "Tag":        nil,
  }

 resp, err := digioceancloud.CreateDns(create)
 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete DNS record

```js
  delete1 := map[string]string{
    "DomainName": "example.com",
    "RecordID":   "28448433",
  }

  resp, err := digioceancloud.DeleteDns(delete1)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### List DNS records

```js
  listRecords := map[string]string{
    "DomainName": "example.com",
  }

  resp, err := digioceancloud.ListDns(listRecords)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
