package hostpool

import (
	"encoding/xml"
	"time"

	"github.com/marthjod/gocart/ocatypes"
	"github.com/marthjod/gocart/vmpool"
)

type HostPool struct {
	XMLName xml.Name `xml:"HOST_POOL"`
	Hosts   []*Host  `xml:"HOST"`
}

type Host struct {
	*ocatypes.Host
	VmPool *vmpool.VmPool
}

func NewHostPool() *HostPool {
	p := new(HostPool)
	return p
}

func (hostPool *HostPool) Read(xmlData []byte) (time.Duration, error) {
	var (
		err     error
		elapsed time.Duration
	)

	_, elapsed, err = ocatypes.Read(xmlData, hostPool)
	return elapsed, err
}

func (hostPool *HostPool) GetHostsInCluster(cluster string) *HostPool {
	var (
		hostpool HostPool
	)
	for _, host := range hostPool.Hosts {
		if host.Cluster == cluster {
			hostpool.Hosts = append(hostpool.Hosts, host)
		}
	}
	return &hostpool
}

func (host *Host) MapVms(vmpool *vmpool.VmPool) {
	host.VmPool = vmpool.GetVmById(host.VmIds...)
}

func (hostPool *HostPool) MapVms(vmpool *vmpool.VmPool) {
	for _, host := range hostPool.Hosts {
		host.MapVms(vmpool)
	}
}
