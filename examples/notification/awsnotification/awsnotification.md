# gocloud compute - AWS

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

### Create Topic

```js
createtopic := map[string]string{
  "Region": "us-east-1",
  "Name": "Bokya",
}

resp, err := amazoncloud.CreateTopic(createtopic)
 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete Topic
```js
deletetopic := map[string]string{
  "Region": "us-east-1",
  "TopicArn" :"arn:aws:sns:us-east-1:478991680879:Bokya",
}

resp , err := amazoncloud.DeleteTopic(deletetopic)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### List Topic

```js

  listtopic := map[string]string{
    "Region": "us-east-1",

  }

  resp, err := amazoncloud.ListTopic(listtopic)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
