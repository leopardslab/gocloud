```
package digiocean

import "github.com/cloudlibz/gocloud/digioceanauth"
```

### TYPES

TokenSource struct for representing DigiOcean credentials.
```
type TokenSource struct {
    DigiOceanAccessToken string
}
```

### VARIABLES

Token is a variable of type TokenSource.
```
var Token TokenSource
```

### FUNCTIONS

LoadConfig loads the DigitalOcean credentials.
```
func LoadConfig()
```
