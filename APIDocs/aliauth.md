```
package aliauth
    import "github.com/cloudlibz/gocloud/aliauth"
```

### FUNCTIONS

```
func LoadConfig()
```

LoadConfig represents Load Ali-cloud config from alicloudconfig.json or environment variables

```
func PutStructToMap(i interface{}) map[string]interface{}
```

PutStructToMap puts key and value of struct into map[string]interface{}

> if value is string and not empty -> put
> if value is bool                 -> put   NOTE: the default value of origin Ali API 's parameter must be false, if not ,do not use this function
> if value is int and not 0        -> put   NOTE: the optional values of origin Ali API 's parameter must dose not include 0, if not ,do not use this function

```
func SignAndDoRequest(action string, params map[string]interface{}, response map[string]interface{}) error
```

SignAndDoRequest sign and do request by action parameter and specific parameters

### TYPES

```
type Configuration struct {
    AliAccessKeyID     string
    AliAccessKeySecret string
}
```

Configuration struct for representing Ali-cloud credentials

```
var Config Configuration
```

Config from alicloudconfig.json
