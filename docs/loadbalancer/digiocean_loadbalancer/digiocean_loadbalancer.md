# gocloud loadbancer - DigitalOcean

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

### Create loadbalancer

```js
  forwardingrules := []map[string]interface{}{
      {
        "EntryProtocol":  "https",
        "EntryPort":  444,
        "TargetProtocol": "https",
        "TargetPort": 443,
        "TLSPassthrough": true,
      },
  }

  healthcheck := map[string]interface{}{
    "Protocol": "http",
    "Port": 80,
    "Path": "/",
    "CheckIntervalSeconds": 10,
    "ResponseTimeoutSeconds": 5,
    "HealthyThreshold": 5,
    "UnhealthyThreshold": 3,
  }

  stickysessions := map[string]interface{}{
    "Type": "none",
  }

  create := map[string]interface{}{
    "Name": "example-01",
    "Algorithm":  "round_robin",
    "Region": "nyc3",
    "ForwardingRules":  forwardingrules,
    "HealthCheck":  healthcheck,
    "StickySessions": stickysessions,
    "DropletIDs":  []int{3164444, 3164445},
    "Tag":  nil,
    "RedirectHTTPToHTTPS": false,
  }

 resp, err := digioceancloud.CreateLoadBalancer(create)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete loadbalancer

```js
  delete := map[string]string{
    "ID": "86407564",
   }

  resp, err := digioceancloud.DeleteLoadBalancer(delete)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Attach node with loadbalancer

```js
  attachnodewithloadbalancer := map[string]interface{}{
    "LoadBalancerID":   "my-load-balancer",
    "DropletIDs":        []int{31331, 31313},
  }

  resp, err := digioceancloud.AttachNodeWithLoadBalancer(attachnodewithloadbalancer)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Detach node with loadbalancer

```js
  detachnodewithloadbalancer := map[string]interface{}{
    "LoadBalancerID":   "my-load-balancer",
    "DropletIDs":        []int{31331, 31313},
  }

 resp, err := digioceancloud.DetachNodeWithLoadBalancer(detachnodewithloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```


### List loadbalancer

```js

 resp, err := digioceancloud.ListLoadBalancer(nil)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
