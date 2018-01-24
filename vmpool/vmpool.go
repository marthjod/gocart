package vmpool

import (
	"encoding/xml"
	"fmt"
	"io"
	"regexp"

	"github.com/marthjod/gocart/ocatypes"
)

type VmPool struct {
	XMLName xml.Name       `xml:"VM_POOL"`
	Vms     []*ocatypes.Vm `xml:"VM"`
}

func (vmpool *VmPool) String() string {
	var list = []string{}

	for _, vm := range vmpool.Vms {
		list = append(list, vm.Name)
	}

	return fmt.Sprintf("%s", list)
}

// ApiMethod implements the api.Endpointer interface
func (vmpool *VmPool) ApiMethod() string {
	return "one.vmpool.info"
}

// ApiArgs implements the api.Endpointer interface
// API parameter documentation: http://docs.opennebula.org/4.10/integration/system_interfaces/api.html#one-vmpool-info
func (vmpool *VmPool) ApiArgs(authstring string) []interface{} {
	return []interface{}{authstring, -2, -1, -1, -1}
}

func (vmpool *VmPool) Unmarshal(data []byte) error {
	err := xml.Unmarshal(data, vmpool)
	return err
}

func NewVmPool() *VmPool {
	p := new(VmPool)
	return p
}

func FromReader(r io.Reader) (*VmPool, error) {
	pool := VmPool{}
	dec := xml.NewDecoder(r)
	if err := dec.Decode(&pool); err != nil {
		return nil, err
	}
	return &pool, nil
}

func (vmPool *VmPool) GetVmsById(ids ...int) *VmPool {
	var (
		pool VmPool
	)
	for _, vm := range vmPool.Vms {
		for _, id := range ids {
			if vm.Id == id {
				pool.Vms = append(pool.Vms, vm)
			}
		}
	}
	return &pool
}

func (vmPool *VmPool) GetVmsByName(matchPattern string) (*VmPool, error) {
	var pool VmPool
	for _, vm := range vmPool.Vms {
		match, err := regexp.MatchString(matchPattern, vm.Name)
		if err != nil {
			return &pool, err
		}
		if match {
			pool.Vms = append(pool.Vms, vm)
		}
	}
	return &pool, nil
}

func (vmPool *VmPool) GetDistinctVmNamePatterns(filter, prefix, infix, suffix string) map[string]bool {
	vmNameExtractor := func(vm *ocatypes.Vm) string {
		return vm.Name
	}

	return vmPool.GetDistinctVmNamePatternsExtractHostname(filter, prefix, infix, suffix, vmNameExtractor)
}

func (vmPool *VmPool) GetDistinctVmNamePatternsExtractHostname(filter, prefix, infix, suffix string,
	hostNameExtractor func(vm *ocatypes.Vm) string) map[string]bool {

	var (
		distinctPatterns = make(map[string]bool, 0)
		pattern          string
	)

	re := regexp.MustCompile(filter)

	for _, vm := range vmPool.Vms {

		groups := re.FindStringSubmatch(hostNameExtractor(vm))
		if groups == nil {
			continue
		}

		if len(groups) >= 3 {
			pattern = fmt.Sprintf("%s%s%s%s%s", prefix, groups[1], infix, groups[2], suffix)
			distinctPatterns[pattern] = true
		} else {
			distinctPatterns[vm.Name] = true
		}
	}

	return distinctPatterns
}

func (vmPool *VmPool) GetVmsByLCMState(state ocatypes.LCMState) (*VmPool, error) {
	var pool VmPool
	for _, vm := range vmPool.Vms {
		if vm.LCMState == state {
			pool.Vms = append(pool.Vms, vm)
		}
	}
	return &pool, nil
}
