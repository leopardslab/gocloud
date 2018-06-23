# gocloud container - Ali-cloud

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

### Create cluster

```
    create := map[string]interface{}{
        "password":           "TestPwd124",
        "region_id":          "cn-beijing",
        "instance_type":      "ecs.n1.small",
        "name":               "test-cluster",
        "size":               1,
        "network_mode":       "vpc",
        "vpc_id":             "vpc-2ze578wokbm1ykyr7d8w6",
        "vswitch_id":         "vsw-2zeryg1zeofqj0u7o6buw",
        "subnet_cidr":        "172.29.1.0/24",
        "data_disk_category": "cloud_ssd",
        "data_disk_size":     40,
        "need_slb":           true,
        "ecs_image_id":       "centos_7_04_64_20G_alibase_201701015",
        "io_optimized":       "true",
        "release_eip_flag":   false,
    }

    resp, err := alicloud.Createcluster(create)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    create, err := alicontainer.NewCreateClusterBuilder().
        Password("TestPwd123").
        RegionID("cn-beijing").
        InstanceType("ecs.n1.small").
        Name("my-go-cloud-test-cluster").
        Size(1).NetworkMode("vpc").
        VPCID("vpc-2ze578wokbm1ykyr7d8w6").
        VSwitchID("vsw-2zeryg1zeofqj0u7o6buw").
        SubnetCIDR("172.26.1.0/24").
        DataDiskCategory("cloud_ssd").
        DataDiskSize(40).NeedSLB(false).
        ECSImageID("centos_7_04_64_20G_alibase_201701015").
        IOOptimized("true").ReleaseEipFlag(false).
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Createcluster(create)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Delete cluster

```
  delete := map[string]interface{}{
          "region_id":  "cn-beijing",
          "cluster_id": "ce82a3a70d4ea4a71826ae877c5ee31f2",
  }

  resp, err := alicloud.Deletecluster(delete)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    deleteCluster, err := alicontainer.NewDeleteClusterBuilder().
        RegionID("cn-beijing").
        ClusterID("c5979e038f7154f28b7e8e9b4e79cd39e").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Deletecluster(deleteCluster)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Run task

```
  alicloud, _ := gocloud.CloudProvider(gocloud.Aliprovider)
  runTask := map[string]interface{}{
          "cluster_id":  "ce82a3a70d4ea4a71826ae877c5ee31f2",
          "name":        "test",
          "description": "This is a test application",
          "template":    "web:\r\n  image: nginx",
          "version":     "1.0",
          "environment": map[string]string{"USER": "abc", "PWD": "password",},
  }

  resp, err := alicloud.Runtask(runTask)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    runTask, err := alicontainer.NewRunTaskBuilder().
        ClusterID("c5446f20698be420e89ec5a2097685faa").
        Name("test-go-cloud").
        Description("This is a test application").
        Template("web:\r\n  image: nginx").
        Version("1.0").
        Environment(map[string]string{"USER": "abc", "PWD": "password"}).
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Runtask(runTask)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
    fmt.Println(response["status"])
```

### Start task

```
  startTask := map[string]interface{}{
          "cluster_id": "ce82a3a70d4ea4a71826ae877c5ee31f2",
          "name":       "test",
  }

  resp, err := alicloud.Starttask(startTask)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    startTask, err := alicontainer.NewStartTaskBuilder().
        ClusterID("c5446f20698be420e89ec5a2097685faa").
        Name("test").
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Starttask(startTask)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
    fmt.Println(response["status"])
```

### Stop task

```
  stopTask := map[string]interface{}{
          "cluster_id": "ce82a3a70d4ea4a71826ae877c5ee31f2",
          "name":       "test",
          "timeout":    20,
  }

  resp, err := alicloud.Stoptask(stopTask)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

or

```
    stopTask, err := alicontainer.NewStopTaskBuilder().
        ClusterID("c5446f20698be420e89ec5a2097685faa").
        Name("test").
        Timeout(20).
        Build()
    if err != nil {
        fmt.Println(err)
        return
    }
    resp, err := alicloud.Stoptask(stopTask)
    if err != nil {
        fmt.Println(err)
        return
    }
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```
