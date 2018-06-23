# gocloud loadbancer - Ali-cloud

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

### Create loadbalancer

```
  createLoadBalancer := map[string]interface{}{
          "RegionId":           "cn-qingdao",
          "LoadBalancerName":   "abc",
          "AddressType":        "internet",
          "InternetChargeType": "paybytraffic",
  }

  resp, err := alicloud.Creatloadbalancer(createLoadBalancer)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    create, err := aliloadbalancer.NewCreateLoadBalancerBuilder().
        RegionID("cn-qingdao").
        LoadBalancerName("abc").
        AddressType("internet").
        InternetChargeType("paybytraffic").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Createloadbalancer(create)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Delete loadbalancer

```
  deleteLoadBalancer := map[string]interface{}{
          "RegionId":       "cn-qingdao",
          "LoadBalancerId": "lb-m5eavuvgjh0pho3hm4ub5",
  }

  resp, err := alicloud.Deleteloadbalancer(deleteLoadBalancer)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    deleteLoadBalancer, err := aliloadbalancer.NewDeleteLoadBalancerBuilder().
        RegionID("cn-qingdao").
        LoadBalancerID("lb-a2dgmwo53mftn7h34za84").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Deleteloadbalancer(deleteLoadBalancer)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### List loadbalancer

```
  list := map[string]interface{}{
          "RegionId":       "cn-qingdao",
  }

  resp, err := alicloud.Listloadbalancer(list)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    list, err := aliloadbalancer.NewListLoadBalancerBuilder().
        RegionID("cn-qingdao").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Listloadbalancer(list)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Attach node with loadbalancer

```
  attach := map[string]interface{}{
          "LoadBalancerId": "lb-m5eemmwmtmt4l6jf73d72",
          "BackendServers": "[{'ServerId':'i-m5ecx5g9m0cflv1k83zu','Weight':'100','Type':'ecs'}," +
              "{'ServerId':'i-m5eahbbwqxawpj1opww9','Weight':'100','Type':'ecs'}]",
  }

  resp, err := alicloud.Attachnodewithloadbalancer(attach)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    attach, err := aliloadbalancer.NewAttachLoadBalancerBuilder().
        LoadBalancerID("lb-m5e7ldi4obgcya3bju3wu").
        BackendServers("[{'ServerId':'i-m5e0cjxe9c7iulv9znvw','Weight':'100','Type':'ecs'}," +
        "{'ServerId':'i-m5efh52hzdkyjoaafwc0','Weight':'100','Type':'ecs'}]").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Attachnodewithloadbalancer(attach)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Detach node with loadbalancer

```
  detach := map[string]interface{}{
          "RegionId":       "cn-qingdao",
          "LoadBalancerId": "lb-m5eemmwmtmt4l6jf73d72",
          "BackendServers": "[{'ServerId':'i-m5ecx5g9m0cflv1k83zu','Type':'ecs'}," +
              "{'ServerId':'i-m5eahbbwqxawpj1opww9','Type':'ecs'}]",
  }

  resp, err := alicloud.Detachnodewithloadbalancer(detach)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    detach, err := aliloadbalancer.NewDetachLoadBalancerBuilder().
        RegionID("cn-qingdao").
        LoadBalancerID("lb-m5eemmwmtmt4l6jf73d72").
        BackendServers("[{'ServerId':'i-m5ecx5g9m0cflv1k83zu','Type':'ecs'}," +
        "{'ServerId':'i-m5eahbbwqxawpj1opww9','Type':'ecs'}]").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Detachnodewithloadbalancer(detach)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```