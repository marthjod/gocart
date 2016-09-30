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
	XMLName    xml.Name `xml:"TEMPLATE"`
	Cpu        string   `xml:"CPU"`
	Disk       string   `xml:"DISK"`
	Memory     string   `xml:"MEMORY"`
	Name       string   `xml:"NAME"`
	Nics       []Nic    `xml:"NIC"`
	VCpu       string   `xml:"VCPU"`
	Datacenter string   `xml:"DATACENTER"`
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
	XMLName   xml.Name `xml:"HOST"`
	Id        int      `xml:"ID"`
	Name      string   `xml:"NAME"`
	Cluster   string   `xml:"CLUSTER"`
	ClusterId int      `xml:"CLUSTER_ID"`
	Template  Template `xml:"TEMPLATE"`
	VmIds     []int    `xml:"VMS>ID"`
}

type VmPool struct {
	XMLName xml.Name `xml:"VM_POOL"`
	Vms     []Vm     `xml:"VM"` // ?
}

func Read(xmlData []byte, pool interface{}) (interface{}, time.Duration, error) {
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
