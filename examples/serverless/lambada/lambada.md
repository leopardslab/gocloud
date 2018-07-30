# gocloud serverless - lambada

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

amazoncloud := gocloud.AmazonProvider()

```

### Create Function

```js
code := map[string]interface{}{
  "ZipFile": "gocloud-ecce2422-f215-4d44-83a8-8361b457c5d9",
}

createfunction := map[string]interface{}{
  "Region":       "us-east-1",
  "FunctionName": "gocloud3",
  "Role":         "arn:aws:iam::478991680879:role/service-role/bokya",
  "Runtime":      "nodejs6.10",
  "Handler":      "index.handler",
  "Code":         code,
}

resp, err := amazoncloud.CreateFunction(createfunction)

response := resp.(map[string]interface{})

fmt.Println(response["body"])
```

### list function

```js

listfunction := map[string]interface{}{
  "Region":     "us-east-1",
}


resp, err := amazoncloud.ListFunction(listfunction)


response := resp.(map[string]interface{})

fmt.Println(response["body"])
```


### list function

```js

getfunction := map[string]interface{}{
  "Region":       "us-east-1",
  "FunctionName": "gocloud",
}

resp, err := amazoncloud.GetFunction(getfunction)

response := resp.(map[string]interface{})

fmt.Println(response["body"])
```


### Delete Function

```js

deletefunction	 := map[string]interface{}{
  "Region":     "us-east-1",
  "FunctionName" : "gocloud2",
}


resp, err := amazoncloud.DeleteFunction(deletefunction)

response := resp.(map[string]interface{})

fmt.Println(response["body"])
```
