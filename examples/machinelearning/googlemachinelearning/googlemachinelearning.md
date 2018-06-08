# gocloud machine learning - ml engine

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

### Create MLModel

```js
defaultVersion := make(map[string]interface{})

defaultVersion["Name"] = "bokkya"
defaultVersion["DeploymentUri"] = "gs:bokkya"

createMLModel := map[string]interface{}{
	"Parent": "projects/adept-comfort-202709",
	"Name"  :"pratik",
	"DefaultVersion" :defaultVersion,
}

resp, err := googlecloud.CreateMLModel(createMLModel)

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])


  ```

### Get MLModel

```js
getMLModel := map[string]string{
	"name": "projects/adept-comfort-202709/models/hello",
}

resp, err := googlecloud.GetMLModel(getMLModel)

if err != nil {
	t.Errorf("Test Fail")
}

response := resp.(map[string]interface{})
fmt.Println(response["body"])
```

###Update MLModel

```js


	defaultVersion := make(map[string]interface{})

	defaultVersion["Name"] = "bokkya"

	updateMLModel := map[string]interface{}{
		"Parent": "projects/adept-comfort-202709/models/pratik",
		"Name"  :"pratik",
		"DefaultVersion" :defaultVersion,
		"UpdateMask" : "description",
	}

	resp, err := googlecloud.UpdateMLModel(updateMLModel)

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])```
