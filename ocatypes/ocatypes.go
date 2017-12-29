package ocatypes

import (
	"encoding/xml"
	"fmt"
)

// ClusterPool is a list of clusters
type ClusterPool struct {
	XMLName  xml.Name   `xml:"CLUSTER_POOL"`
	Clusters []*Cluster `xml:"CLUSTER"`
}

// Cluster represents a cluster
type Cluster struct {
	XMLName      xml.Name `xml:"CLUSTER"`
	Name         string   `xml:"NAME"`
	ID           int      `xml:"ID"`
	DatastoreIDs []int    `xml:"DATASTORES"`
	VnetIDs      []int    `xml:"VNETS"`
}

// DSPool is a list of Datastores
type DSPool struct {
	XMLName    xml.Name     `xml:"DATASTORE_POOL"`
	Datastores []*Datastore `xml:"DATASTORE"`
}

// Datastore
type Datastore struct {
	XMLName   xml.Name `xml:"DATASTORE"`
	Name      string   `xml:"NAME"`
	ID        int      `xml:"ID"`
	ClusterID int      `xml:"CLUSTER_ID"`
	Cluster   string   `xml:"CLUSTER"`
}

// VNetPool is a list of Virtual Networks
type VNetPool struct {
	XMLName  xml.Name `xml:"VNET_POOL"`
	Networks []*VNet  `xml:"VNET"`
}

// VNet represents a virtual network
type VNet struct {
	XMLName   xml.Name `xml:"VNET"`
	Name      string   `xml:"NAME"`
	ID        int      `xml:"ID"`
	Cluster   string   `xml:"CLUSTER"`
	ClusterID int      `xml:"CLUSTER_ID"`
	Bridge    string   `xml:"BRIDGE"`
}

// Disk represents a disk
type Disk struct {
	XMLName xml.Name `xml:"DISK"`
	Name    string   `xml:"IMAGE"`
	ID      int      `xml:"IMAGE_ID"`
}

// Image represents an image
type Image struct {
	XMLName     xml.Name `xml:"IMAGE"`
	ID          int      `xml:"ID"`
	Name        string   `xml:"NAME"`
	Datastore   string   `xml:"DATASTORE"`
	DatastoreID int      `xml:"DATASTORE_ID"`
	RunningVMS  int      `xml:"RUNNING_VMS"`
}

// Nic represents an network interface
type Nic struct {
	XMLName   xml.Name `xml:"NIC"`
	Name      string   `xml:"NETWORK"`
	NetworkId int      `xml:"NETWORK_ID"`
}

// VMTemplatePool is a list of VMTemplates
type VMTemplatePool struct {
	XMLName   xml.Name      `xml:"VMTEMPLATE_POOL"`
	Templates []*VmTemplate `xml:"VMTEMPLATE"`
}

type HostTemplate struct {
	XMLName        xml.Name `xml:"TEMPLATE"`
	Cpu            string   `xml:"CPU"`
	Disk           []Disk   `xml:"DISK"`
	Memory         string   `xml:"MEMORY"`
	Name           string   `xml:"NAME"`
	Nics           []Nic    `xml:"NIC"`
	VCpu           string   `xml:"VCPU"`
	Datacenter     string   `xml:"DATACENTER"`
	Requirements   string   `xml:"REQUIREMENTS"`
	DSRequirements string   `xml:"SCHED_DS_REQUIREMENTS"`
	Items          Tags     `xml:",any"`
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
	Disk     []Disk       `xml:"DISK"`
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
