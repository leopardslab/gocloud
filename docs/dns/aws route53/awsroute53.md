# gocloud DNS - AWS

## Configure AWS credentials

Create `gocloudconfig.json` as follows,
```js
{
  "AWSAccessKeyID": "xxxxxxxxxxxx",
  "AWSSecretKey": "xxxxxxxxxxxx",
}
```

also You can setup enviroment variables as

```js
export AWSAccessKeyID =  "xxxxxxxxxxxx"
export AWSSecretKey = "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/shlokgilda/gocloud/gocloud"

amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
```

### create dns

```js
  createdns := map[string]interface{}{
		"name":             "rootmonk.me",
		"hostedZoneConfig": "hostedZoneConfig",
	}

 resp, err := awsdns.Createdns(createdns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### delete dns
```js
 deletedns := map[string]string{
		"ID": "ZOD7SUP0ZGGQQ",
	}

 resp, err := awsdns.Deletedns(deletedns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### List dns

```js
 listdns := map[string]interface{}{
		"marker":   "",
		"maxItems": 2,
	}

  resp, err := awsdns.Listdns(listdns)
 
 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### list ResourcednsRecordSets

```js
 
 listResourcednsRecordSets := map[string]interface{}{
	"zone": "ZBNX5TIW033J2",
  }

 resp, err := awsdns.ListResourcednsRecordSets(listResourcednsRecordSets)
 
 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
