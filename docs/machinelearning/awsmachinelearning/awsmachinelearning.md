# gocloud machinelearning - awsmachinelearning

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

import "github.com/cloudlibz/gocloud/gocloud"

amazoncloud := gocloud.AmazonProvider()
```

### CreateMLModel

```js
createMLModel := map[string]interface{}{
		"Region":      "us-east-1",
		"MLModelName": "EXAMPLE",
		"MLModelId":   "ml-EL5FRUNlk7p",
		"MLModelType": "REGRESSION",
		"RecipeUri": "	s3://bokya/census.csv",
		"TrainingDataSourceId": "ds-Lf3D4KaPukx",
	}

	resp, err := amazoncloud.CreateMLModel(createMLModel)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Update MLModel

```js
updateMLModel := map[string]interface{}{
		"Region":    "us-east-1",
		"MLModelId": "ds-Lf3D4KaPukx",
	}

	resp, err := amazoncloud.UpdateMLModel(updateMLModel)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

###Delete MLModel

```js
deleteMLModel := map[string]interface{}{
		"Region":    "us-east-1",
		"MLModelId": "ds-Lf3D4KaPukx",
	}

	resp, err := amazoncloud.DeleteMLModel(deleteMLModel)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Get MLModel

```js
  getMLModel := map[string]interface{}{
    "Region":    "us-east-1",
    "MLModelId": "ml-EL5FRUNlk7p",
  }

  resp, err := amazoncloud.GetMLModel(getMLModel)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
