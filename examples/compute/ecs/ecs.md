# gocloud compute - ECS

## Configure Ali-cloud credentials

Create `alicloudconfig.json` as follows,
```js
{
  "AliAccessKeyID": "xxxxxxxxxxxx",
  "AliAccessKeySecret": "xxxxxxxxxxxx"
}
```

also You can setup enviroment variables as

```js
export AliAccessKeyID =  "xxxxxxxxxxxx"
export AliAccessKeySecret = "xxxxxxxxxxxx"
```

## Initialize library

```js
import "github.com/cloudlibz/gocloud/gocloud"

alicloud, _ := gocloud.CloudProvider(gocloud.Aliprovider)
```

### Create instance

```js
  create := map[string]interface{}{
        "RegionId":        "cn-qingdao",
        "ImageId":         "centos_7_04_64_20G_alibase_201701015.vhd",
	    "InstanceType":    "ecs.xn4.small",
	    "SecurityGroupId": "sg-m5egbo9s5xb21kpu6nk2",
  }

  resp, err := alicloud.Createnode(create)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Start instance

```js
  start := map[string]interface{}{
        "InstanceId": "i-m5e3ee3z8wdy8ktdq591",
  }

  resp, err := alicloud.Startnode(start)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Stop instance

```js
  stop := map[string]interface{}{
        "InstanceId": "i-m5e3ee3z8wdy8ktdq591",
		"ForceStop":  false,
  }

  resp, err := alicloud.Stopnode(stop)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Reboot instance

```js
  reboot := map[string]interface{}{
		"InstanceId": "i-m5e3ee3z8wdy8ktdq591",
		"ForceStop":  false,
  }

  resp, err := alicloud.Rebootnode(reboot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete instance

```js
  dlete := map[string]interface{}{
		"InstanceId": "i-m5e3ee3z8wdy8ktdq591",
  }

  resp, err := alicloud.Deletenode(delete)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
