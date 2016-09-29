package vmpool

import (
	"encoding/xml"
	"regexp"
	"time"

	"github.com/marthjod/gocart/ocatypes"
)

type VmPool struct {
	XMLName xml.Name `xml:"VM_POOL"`
	Vms     []*Vm    `xml:"VM"`
}

type Vm struct {
	ocatypes.Vm
}

func (vm Vm) String() string {
	return vm.Name
}

func NewVmPool() *VmPool {
	p := new(VmPool)
	return p
}

func (vmPool *VmPool) Read(xmlData []byte) (time.Duration, error) {
	var (
		err     error
		elapsed time.Duration
	)

	_, elapsed, err = ocatypes.Read(xmlData, vmPool)
	return elapsed, err
}

func (vmPool *VmPool) GetVmById(ids ...int) *VmPool {
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

func (vmPool *VmPool) GetVmByName(matchPattern string) (*VmPool, error) {
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
