# gocloud kinesis - stream data processing

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

### List Stream

```js
liststream := map[string]interface{}{
  "Region":    "us-east-1",
}

resp, err := amazoncloud.ListStream(liststream)

if err != nil {
  t.Errorf("Test Fail")
}

response := resp.(map[string]interface{})

fmt.Println(response["body"])
```

### Create Stream

```js
createstream := map[string]interface{}{
  "Region":    "us-east-1",
  "StreamName" : "gocloud",
  "ShardCount" : 1,
}

resp, err := amazoncloud.CreateStream(createstream)

if err != nil {
  t.Errorf("Test Fail")
}

response := resp.(map[string]interface{})

fmt.Println(response["body"])
```

###Describe Stream

```js
describestream := map[string]interface{}{
  "Region":    "us-east-1",
  "StreamName" : "gocloud",
}

resp, err := amazoncloud.DescribeStream(describestream)

if err != nil {
  t.Errorf("Test Fail")
}

response := resp.(map[string]interface{})

fmt.Println(response["body"])
```

### Delete Stream

```js
deletestream := map[string]interface{}{
  "Region":    "us-east-1",
  "StreamName" : "gocloud",
}

resp, err := amazoncloud.DeleteStream(deletestream)

if err != nil {
  t.Errorf("Test Fail")
}

response := resp.(map[string]interface{})

fmt.Println(response["body"])
```
