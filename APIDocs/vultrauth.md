```
package vultrauth
    import "github.com/cloudlibz/gocloud/vultrauth"
```

### FUNCTIONS

```
func LoadConfig()
```
LoadConfig represents Load Vultr config from vultrconfig.json or environment variables

```
func SignAndDoRequest(method string, path string, params map[string]interface{}, response map[string]interface{}) error
```
SignAndDoRequest sign and do request by action parameter and specific parameters


### TYPES

```
type Configuration struct {
    VultrAPIKey string
}
```
Configuration struct for representing Vultr API Key

```
var Config Configuration
```
Config from vultrconfig.json

