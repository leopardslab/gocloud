# gocloud serverless - google cloud functions

## Configure Google Cloud credentials.

Download your service account credentials file from Google Cloud and save it as `googlecloudinfo.json` in your <b>HOME</b> directory.

You can also set the credentials as environment variables:
```js
export PrivateKey =  "xxxxxxxxxxxx"
export Type =  "xxxxxxxxxxxx"
export ProjectID = "xxxxxxxxxxxx"
export PrivateKeyID = "xxxxxxxxxxxx"
export ClientEmail = "xxxxxxxxxxxx"
export ClientID = "xxxxxxxxxxxx" .
export AuthURI = "xxxxxxxxxxxx"
export TokenURI = "xxxxxxxxxxxx"
export AuthProviderX509CertURL = "xxxxxxxxxxxx"
export ClientX509CertURL =  "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
```

### Create function

```js
httpsTrigger := make(map[string]string)
httpsTrigger["URL"] = "https://us-central1-adept-comfort-202709.cloudfunctions.net/function-1"

labels := make(map[string]string)
labels["deployment-tool"] = "console-cloud"

createfunction := map[string]interface{}{
	"Location":            "projects/adept-comfort-202709/locations/us-central1",
	"Name":                "projects/adept-comfort-202709/locations/us-central1/functions/function-2",
	"Status":              "ACTIVE",
	"HTTPSTrigger":        httpsTrigger,
	"EntryPoint":          "helloWorld",
	"Timeout":             "60s",
	"AvailableMemoryMb":   256,
	"ServiceAccountEmail": "adept-comfort-202709@appspot.gserviceaccount.com",
	"UpdateTime":          "2018-05-11T18:20:33Z",
	"Runtime":             "nodejs6",
	"SourceUploadURL":     "urllink",
	"VersionID":           "1",
	"Labels":              labels,
}

resp, err := googlecloud.Createfunction(createfunction)

response := resp.(map[string]interface{})

fmt.Println(response["body"])

  ```

### List function

```js
listfunction := map[string]string{
	"name":     "projects/adept-comfort-202709/locations/us-central1",
	"pageSize": "1",
}

resp, err := googlecloud.Listfunction(listfunction)

response := resp.(map[string]interface{})

fmt.Println(response["body"])

```

### Delete function

```js

deletefunction := map[string]string{
	"name":"projects/adept-comfort-202709/locations/us-central1/functions/function-1",
}

 resp, err := googlecloud.Deletefunction(deletefunction)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Call function

```js
callfunction := map[string]string{
	"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
}

resp, err := googlecloud.Callfunction(callfunction)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```



### Get function

```js
getfunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
	}

resp, err := googlecloudfunctions.Getfunction(getfunction)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
