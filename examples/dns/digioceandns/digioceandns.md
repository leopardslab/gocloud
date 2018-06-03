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

 resp, err := digioceancloud.Createdns(create)
 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete instance

```js
  delete1 := map[string]string{
    "DomainName": "example.com",
    "RecordID":   "28448433",
  }

  resp, err := digioceancloud.Deletedns(delete1)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
