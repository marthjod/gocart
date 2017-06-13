package ocatypes

import (
	"encoding/xml"
	"fmt"
)

type Nic struct {
	XMLName   xml.Name `xml:"NIC"`
	NetworkId int      `xml:"NETWORK_ID"`
}

type HostTemplate struct {
	XMLName    xml.Name `xml:"TEMPLATE"`
	Cpu        string   `xml:"CPU"`
	Disk       string   `xml:"DISK"`
	Memory     string   `xml:"MEMORY"`
	Name       string   `xml:"NAME"`
	Nics       []Nic    `xml:"NIC"`
	VCpu       string   `xml:"VCPU"`
	Datacenter string   `xml:"DATACENTER"`
	Items      Tags     `xml:",any"`
}

type VmTemplate struct {
	XMLName  xml.Name     `xml:"VMTEMPLATE"`
	Id       int          `xml:"ID"`
	Name     string       `xml:"NAME"`
	Uname    string       `xml:"UNAME"`
	RegTime  int          `xml:"REGTIME"`
	Template HostTemplate `xml:"TEMPLATE"`
	Memory   int          `xml:"MEMORY"`
	VmId     int          `xml:"VMID"`
	Disk     string       `xml:"DISK"`
	Cpu      string       `xml:"CPU"`
}

type Vm struct {
	XMLName      xml.Name     `xml:"VM"`
	Id           int          `xml:"ID"`
	Name         string       `xml:"NAME"`
	Cpu          int          `xml:"CPU"`
	LastPoll     int          `xml:"LAST_POLL"`
	LCMState     int          `xml:"LCM_STATE"`
	Resched      int          `xml:"RESCHED"`
	DeployId     string       `xml:"DEPLOY_ID"`
	Template     HostTemplate `xml:"TEMPLATE"`
	UserTemplate UserTemplate `xml:"USER_TEMPLATE"`
}

type UserTemplate struct {
	Items Tags `xml:",any"`
}

type Tag struct {
	XMLName xml.Name
	Content string `xml:",chardata"`
}

type Tags []Tag

func (tags Tags) GetCustom(tagName string) (string, error) {
	for _, tag := range tags {
		if tagName == tag.XMLName.Local {
			return tag.Content, nil
		}
	}
	return "", fmt.Errorf("tag %s not found", tagName)
}

type Host struct {
	XMLName   xml.Name     `xml:"HOST"`
	Id        int          `xml:"ID"`
	Name      string       `xml:"NAME"`
	State     int          `xml:"STATE"`
	Cluster   string       `xml:"CLUSTER"`
	ClusterId int          `xml:"CLUSTER_ID"`
	Template  HostTemplate `xml:"TEMPLATE"`
	VmIds     []int        `xml:"VMS>ID"`
}

func (h *Host) IsEmpty() bool {
	return len(h.VmIds) == 0
}

type VmPool struct {
	XMLName xml.Name `xml:"VM_POOL"`
	Vms     []Vm     `xml:"VM"` // ?
}
