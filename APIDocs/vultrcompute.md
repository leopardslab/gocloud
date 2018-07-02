```
package vultrcompute
    import "github.com/cloudlibz/gocloud/compute/vultrcompute"
```

### TYPES

```
type VultrCompute struct {
}
```
VultrCompute struct

```
func (vultrCompute *VultrCompute) CreateNode(request interface{}) (resp interface{}, err error)
```
CreateNode function creates a new VultrCompute instance.

```
func (vultrCompute *VultrCompute) DeleteNode(request interface{}) (resp interface{}, err error)
```
DeleteNode function deletes a VultrCompute instance.

```
func (vultrCompute *VultrCompute) ListNode() (resp interface{}, err error)
```
ListNode function lists VultrCompute instances.

```
func (vultrCompute *VultrCompute) RebootNode(request interface{}) (resp interface{}, err error)
```
RebootNode function reboots a VultrCompute instance.

```
func (vultrCompute *VultrCompute) StartNode(request interface{}) (resp interface{}, err error)
```
StartNode function starts a VultrCompute instance.

```
func (vultrCompute *VultrCompute) StopNode(request interface{}) (resp interface{}, err error)
```
StopNode function stops a VultrCompute instance.
