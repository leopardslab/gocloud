# gocloud database - bigtable

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

googlecloud := gocloud.GoogleProvider()

```
### Create tables

```js
	table := make(map[string]interface{})

	initialSplits := make([]map[string]interface{},0)

	createtables := map[string]interface{}{
		"parent": "projects/adept-comfort-202709/instances/helloo",
		"tableId" :"tableId",
		"table"  : table,
		"initialSplits" : initialSplits,
	}

	resp, err := googlecloud.CreateTables(createtables)

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])


  ```

### Describe tables

```js
	describetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	resp, err := googlecloud.DescribeTables(describetables)

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

```

### Delete tables

```js

deletetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	resp, err := googlecloud.DeleteTables(deletetables)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### List tables

```js
listtables := map[string]string{
		"parent":    "projects/adept-comfort-202709/instances/helloo",
		"view":      "NAME_ONLY",
		"pageToken": "",
	}

	resp, err := googlecloud.ListTables(listtables)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```
