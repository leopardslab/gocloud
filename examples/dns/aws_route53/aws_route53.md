# gocloud DNS - AWS

## Configure AWS credentials.

Create `amazoncloudconfig.json` in your <b>HOME/.gocloud</b> directory as follows:
```js
{
  "AWSAccessKeyID": "xxxxxxxxxxxx",
  "AWSSecretKey": "xxxxxxxxxxxx"
}
```

You can also set the credentials as environment variables:
```js
export AWSAccessKeyID =  "xxxxxxxxxxxx"
export AWSSecretKey = "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
```

### Create DNS

```js
  createdns := map[string]interface{}{
		"name":             "rootmonk.me",
		"hostedZoneConfig": "hostedZoneConfig",
	}

 resp, err := awsdns.Createdns(createdns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete DNS
```js
 deletedns := map[string]string{
		"ID": "ZOD7SUP0ZGGQQ",
	}

 resp, err := awsdns.Deletedns(deletedns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### List DNS

```js
 listdns := map[string]interface{}{
		"marker":   "",
		"maxItems": 2,
	}

  resp, err := awsdns.Listdns(listdns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### List ResourcednsRecordSets

```js

 listResourcednsRecordSets := map[string]interface{}{
	"zone": "ZBNX5TIW033J2",
  }

 resp, err := awsdns.ListResourcednsRecordSets(listResourcednsRecordSets)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```