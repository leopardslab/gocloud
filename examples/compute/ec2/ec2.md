# gocloud compute - AWS

## Configure AWS credentials

Create `amazoncloudconfig.json` as follows,
```js
{
  "AWSAccessKeyID": "xxxxxxxxxxxx",
  "AWSSecretKey": "xxxxxxxxxxxx"
}
```

also You can setup enviroment variables as

```js
export AWSAccessKeyID =  "xxxxxxxxxxxx"
export AWSSecretKey = "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
```

### Create instance

```js
  create := map[string]interface{}{
	"ImageId":      "ami-ccf405a5",
	"InstanceType": "t1.micro",
	"Region":       "us-east-1",
  }

 resp, err := amazoncloud.Createnode(create)
 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Stop instance

```js
  stop := map[string]string{
		"instance-id": "i-06d518ba15b68685c",
		"Region":      "us-east-1",
	}

  resp, err := amazoncloud.Stopnode(stop)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Start instance

```js
  start := map[string]string{
		"instance-id": "i-0174bd6f54178e89b",
		"Region":      "us-east-1",
	}

  resp, err := amazoncloud.Startnode(start)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Reboot instance

```js
  Reboot := map[string]string{
		"instance-id": "i-037a9fae81c33ac30",
		"Region":      "us-east-1",
	}


  resp, err := amazoncloud.Rebootnode(Reboot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### delete instance

```js
  delete := map[string]string{
  "instance-id": "i-0174bd6f54178e89b",
  "Region":      "us-east-1",
   }


  resp, err := amazoncloud.Deletenode(delete)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
