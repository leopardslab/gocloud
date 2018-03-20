# gocloud container - AWS

## Configure AWS credentials.

Create `amazoncloudconfig.json` in your <b>HOME</b> directory as follows:
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

import "github.com/shlokgilda/gocloud/gocloud"

amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
```

### Create cluster

```js
  createcluster := map[string]interface{}{
		"clusterName": "gocloud-test",
		"Region":      "us-east-1",
	}

 resp, err := ecscontainer.Createcluster(createcluster)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete cluster

```js
 deletecluster := map[string]interface{}{
		"clusterName": "gocloud-test",
		"Region":      "us-east-1",
	}

  resp, err := ecscontainer.Deletecluster(deletecluster)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Create service

```js
 LoadBalancers := []map[string]interface{}{{
	"containerName":    "rootmonk",
	"loadBalancerName": "us-east-2",
	},
    {
	"containerName":    "rootmonk",
	"loadBalancerName": "us-east-2",
     },
  }

  createservice := map[string]interface{}{
		"serviceName":    "ecs-simple-service",
		"taskDefinition": "ecs-demo",
		"desiredCount":   1,
		"Region":         "us-east-1",
		"LoadBalancers":  LoadBalancers,
	}

  resp, err := ecscontainer.Createservice(createservice)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete service

```js
  deleteservice := map[string]interface{}{
	"cluster": "cluster",
	"service": "service",
	"Region":  "us-east-1",
   }

  resp, err := ecscontainer.Deleteservice(deleteservice)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```