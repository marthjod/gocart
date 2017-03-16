package hostpool

import (
	"encoding/xml"
	"io"

	"github.com/marthjod/gocart/ocatypes"
	"github.com/marthjod/gocart/vmpool"
)

const (
	INIT = iota
	MONITORING_MONITORED  	// Currently monitoring, previously MONITORED
	MONITORED
	ERROR
	DISABLED
	MONITORING_ERROR	// Currently monitoring, previously ERROR
	MONITORING_INIT  	// Currently monitoring, previously initialized
	MONITORING_DISABLED 	// Currently monitoring, previously DISABLED
)

type HostPool struct {
	XMLName xml.Name `xml:"HOST_POOL"`
	Hosts   []*Host  `xml:"HOST"`
}

// ApiMethod implements the api.Endpointer interface
func (hostpool *HostPool) ApiMethod() string {
	return "one.hostpool.info"
}

// ApiArgs implements the api.Endpointer interface
// API parameter documentation: http://docs.opennebula.org/4.10/integration/system_interfaces/api.html#one-hostpool-info
func (hostpool *HostPool) ApiArgs(authstring string) []interface{} {
	return []interface{}{authstring}
}

func (hostpool *HostPool) Unmarshal(data []byte) error {
	err := xml.Unmarshal(data, hostpool)
	return err
}

type Host struct {
	*ocatypes.Host
	VmPool *vmpool.VmPool
}

func (h *Host) String() string {
	return h.Name
}

func NewHostPool() *HostPool {
	p := new(HostPool)
	return p
}

func FromReader(r io.Reader) (*HostPool, error) {
	pool := HostPool{}
	dec := xml.NewDecoder(r)
	if err := dec.Decode(&pool); err != nil {
		return nil, err
	}
	return &pool, nil
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
	host.VmPool = vmpool.GetVmsById(host.VmIds...)
}

func (hostPool *HostPool) MapVms(vmpool *vmpool.VmPool) {
	for _, host := range hostPool.Hosts {
		host.MapVms(vmpool)
	}
}
