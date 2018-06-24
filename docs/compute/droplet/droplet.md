# gocloud compute - DigitalOcean

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

### Create instance

```js
  image := map[string]interface{}{
    "Slug": "ubuntu-16-04-x64",
  }

  create := map[string]interface{}{
    "Name":   "example.com",
    "Region": "nyc3",
    "Size":   "s-1vcpu-1gb",
    "Image": image,
    "SSHKeys": nil,
    "Backups": false,
    "IPv6": true,
    "UserData": nil,
    "PrivateNetworking": nil,
    "Volumes": nil,
    "Monitoring": false,
    "Tags": []string{"web"},
  }

 resp, err := digioceancloud.CreateNode(create)
 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Stop instance

```js
  stop := map[string]string{
    "ID": "86407564",
  }

  resp, err := digioceancloud.StopNode(stop)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Start instance

```js
  start := map[string]string{
    "ID": "86407564",
  }

  resp, err := digioceancloud.StartNode(start)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Reboot instance

```js
  reboot := map[string]string{
    "ID": "86407564",
  }

  resp, err := digioceancloud.RebootNode(reboot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete instance

```js
  delete := map[string]string{
    "ID": "86406952",
  }

  resp, err := digioceancloud.DeleteNode(delete)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
