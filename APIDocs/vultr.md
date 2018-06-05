```
package vultr
    import "github.com/cloudlibz/gocloud/vultr"
```

### TYPES

```
type Vultr struct {
    vultrcompute.VultrCompute
    vultrstorage.VultrStorage
    vultrloadbalancer.VultrLoadBalancer
    vultrcontainer.VultrContainer
    vultrdns.VultrDNS
}
```
Vultr struct represents Vultr cloud provider.

