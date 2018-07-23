# gocloud  google notification - gce

## Configure Google Cloud credentials.

Download your service account credentials file from Google Cloud and save it as `googlecloudinfo.json` in your <b>HOME/.gocloud</b> directory.

You can also set the credentials as environment variables:
```js
export PrivateKey =  "xxxxxxxxxxxx"
export Type =  "xxxxxxxxxxxx"
export ProjectID = "xxxxxxxxxxxx"
export PrivateKeyID = "xxxxxxxxxxxx"
export ClientEmail = "xxxxxxxxxxxx"
export ClientID = "xxxxxxxxxxxx"
export AuthURI = "xxxxxxxxxxxx"
export TokenURI = "xxxxxxxxxxxx"
export AuthProviderX509CertURL = "xxxxxxxxxxxx"
export ClientX509CertURL =  "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

googlecloud := gocloud.GoogleProvider()

```

### List Topic

```js
	listtopic := map[string]string{
		"Project": "gocloud-206919",
		"PageSize" : "",
		"PageToken": "",
	}

	resp, err := googlecloud.ListTopic(listtopic)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Get Topic

```js
	gettopic := map[string]string{
		"Project": "gocloud-206919",
		"Topic" : "gocloud",
	}

	resp, err := googlecloud.GetTopic(gettopic)

 	response := resp.(map[string]interface{})
 	fmt.Println(response["body"])
```

### Delete Topic

```js

deletetopic := map[string]string{
	"Project": "gocloud-206919",
	"Topic" : "gocloud",
}

resp, err := googlecloud.DeleteTopic(deletetopic)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Create Topic

```js
	createtopic := map[string]string{
		"Project": "gocloud-206919",
		"Topic" : "gocloud3",
	}

_, err := googlecloud.CreateTopic(createtopic)
 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
