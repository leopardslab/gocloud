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