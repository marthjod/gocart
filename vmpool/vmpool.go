package vmpool

import (
	"encoding/xml"
	"io"
	"regexp"

	"github.com/marthjod/gocart/ocatypes"
)

type VmPool struct {
	XMLName xml.Name       `xml:"VM_POOL"`
	Vms     []*ocatypes.Vm `xml:"VM"`
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
