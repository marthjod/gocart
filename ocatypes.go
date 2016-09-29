package ocatypes

import (
	"encoding/xml"
	"time"
)

type Nic struct {
	XMLName   xml.Name `xml:"NIC"`
	NetworkId int      `xml:"NETWORK_ID"`
}

type Template struct {
	XMLName xml.Name `xml:"TEMPLATE"`
	Cpu     string   `xml:"CPU"`
	Disk    string   `xml:"DISK"`
	Memory  string   `xml:"MEMORY"`
	Name    string   `xml:"NAME"`
	Nics    []Nic    `xml:"NIC"`
	VCpu    string   `xml:"VCPU"`
	Datacenter string `xml:"DATACENTER"`
}

type VmTemplate struct {
	XMLName  xml.Name `xml:"VMTEMPLATE"`
	Id       int      `xml:"ID"`
	Name     string   `xml:"NAME"`
	Uname    string   `xml:"UNAME"`
	RegTime  int      `xml:"REGTIME"`
	Template Template `xml:"TEMPLATE"`
	Memory   int      `xml:"MEMORY"`
	VmId     int      `xml:"VMID"`
	Disk     string   `xml:"DISK"`
	Cpu      string   `xml:"CPU"`
}

type Vm struct {
	XMLName  xml.Name `xml:"VM"`
	Id       int      `xml:"ID"`
	Name     string   `xml:"NAME"`
	Cpu      int      `xml:"CPU"`
	LastPoll int      `xml:"LAST_POLL"`
	LCMState int      `xml:"LCM_STATE"`
	Resched  int      `xml:"RESCHED"`
	DeployId string   `xml:"DEPLOY_ID"`
	Template Template `xml:"TEMPLATE"`
}

type Host struct {
	XMLName xml.Name `xml:"HOST"`
	Id int `xml:"ID"`
	Template Template `xml:"TEMPLATE"`
}

type VmPool struct {
	XMLName xml.Name `xml:"VM_POOL"`
	Vms     []Vm     `xml:"VM"` // ?
}

type HostPool struct {
	XMLName xml.Name `xml:"HOST_POOL"`
	Hosts []Host `xml:"HOST"`
}

func NewVmPool() *VmPool {
	p := new(VmPool)
	return p
}

func NewHostPool() *HostPool {
	p := new(HostPool)
	return p
}

func read(xmlData []byte, pool interface{}) (interface{}, time.Duration, error) {
	var (
		err     error
		start   time.Time
		elapsed time.Duration
	)

	start = time.Now()
	err = xml.Unmarshal(xmlData, &pool)
	elapsed = time.Since(start)

	return pool, elapsed, err
}

func (vmPool *VmPool) Read(xmlData []byte) (time.Duration, error) {
	var (
		err error
		elapsed time.Duration
	)

	_, elapsed, err = read(xmlData, vmPool)
	return elapsed, err
}

func (hostPool *HostPool) Read(xmlData []byte) (time.Duration, error) {
	var (
		err     error
		elapsed time.Duration
	)

	_, elapsed, err = read(xmlData, hostPool)
	return elapsed, err
}
