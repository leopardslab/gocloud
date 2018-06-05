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
func (vultrCompute *VultrCompute) Createnode(request interface{}) (resp interface{}, err error)
```
Createnode function creates a new VultrCompute instance.

```
func (vultrCompute *VultrCompute) Deletenode(request interface{}) (resp interface{}, err error)
```
Deletenode function deletes a VultrCompute instance.

```
func (vultrCompute *VultrCompute) ListNode() (resp interface{}, err error)
```
ListNode function lists VultrCompute instances.

```
func (vultrCompute *VultrCompute) Rebootnode(request interface{}) (resp interface{}, err error)
```
Rebootnode function reboots a VultrCompute instance.

```
func (vultrCompute *VultrCompute) Startnode(request interface{}) (resp interface{}, err error)
```
Startnode function starts a VultrCompute instance.

```
func (vultrCompute *VultrCompute) Stopnode(request interface{}) (resp interface{}, err error)
```
Stopnode function stops a VultrCompute instance.
