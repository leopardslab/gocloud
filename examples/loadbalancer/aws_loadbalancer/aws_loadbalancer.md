# gocloud loadbancer - AWS

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

### Create loadbalancer

```js
 Listeners := []map[string]string{{
	"InstancePort":     "80",
	"LoadBalancerPort": "80",
	"Protocol":         "http",
	"InstanceProtocol": "http",
	"SSLCertificateId": "",
	},
 }

 Subnets := []string{"subnet-b59a4599", "subnet-32f9727a"}

 creatloadbalancer := map[string]interface{}{
	"LoadBalancerName": "my-load-balancer",
	"Listeners":        Listeners,
	"Subnets":          Subnets,
 }

 resp, err := awsloadbalancer.CreateLoadBalancer(creatloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete loadbalancer

```js
  deleteloadbalancer := map[string]string{
		"LoadBalancerName": "my-load-balancer",
	}

  resp, err := awsloadbalancer.DeleteLoadBalancer(deleteloadbalancer)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Attach node with loadbalancer

```js
  attachnodewithloadbalancer := map[string]interface{}{
	"Instances":        []string{"i-05f4f2535c41b680b"},
	"LoadBalancerName": "my-load-balancer",
  }

  resp, err := awsloadbalancer.AttachNodeWithLoadBalancer(attachnodewithloadbalancer)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Detach node with loadbalancer

```js
 detachnodewithloadbalancer := map[string]interface{}{
		"Instances":        []string{"i-05f4f2535c41b680b"},
		"LoadBalancerName": "my-load-balancer",
	}

 resp, err := awsloadbalancer.DetachNodeWithLoadBalancer(detachnodewithloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```


### List loadbalancer

```js

 resp, err := awsloadbalancer.ListLoadBalancer(nil)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
