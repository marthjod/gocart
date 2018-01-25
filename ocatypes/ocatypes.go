package ocatypes

import (
	"encoding/xml"
	"fmt"
)

// LCMState represents LCM (lifecycle manager) state
type LCMState int

//go:generate stringer -type=LCMState
const (
	LcmInit                      LCMState = iota
	Prolog                       LCMState = iota
	Boot                         LCMState = iota
	Running                      LCMState = iota
	Migrate                      LCMState = iota
	SaveStop                     LCMState = iota
	SaveSuspend                  LCMState = iota
	SaveMigrate                  LCMState = iota
	PrologMigrate                LCMState = iota
	PrologResume                 LCMState = iota
	EpilogStop                   LCMState = iota
	Epilog                       LCMState = iota
	Shutdown                     LCMState = iota
	CleanupResubmit              LCMState = iota
	Unknown                      LCMState = iota
	Hotplug                      LCMState = iota
	ShutdownPoweroff             LCMState = iota
	BootUnknown                  LCMState = iota
	BootPoweroff                 LCMState = iota
	BootSuspended                LCMState = iota
	BootStopped                  LCMState = iota
	CleanupDelete                LCMState = iota
	HotplugSnapshot              LCMState = iota
	HotplugNic                   LCMState = iota
	HotplugSaveas                LCMState = iota
	HotplugSaveasPoweroff        LCMState = iota
	HotplugSaveasSuspended       LCMState = iota
	ShutdownUndeploy             LCMState = iota
	EpilogUndeploy               LCMState = iota
	PrologUndeploy               LCMState = iota
	BootUndeploy                 LCMState = iota
	HotplugPrologPoweroff        LCMState = iota
	HotplugEpilogPoweroff        LCMState = iota
	BootMigrate                  LCMState = iota
	BootFailure                  LCMState = iota
	BootMigrateFailure           LCMState = iota
	PrologMigrateFailure         LCMState = iota
	PrologFailure                LCMState = iota
	EpilogFailure                LCMState = iota
	EpilogStopFailure            LCMState = iota
	EpilogUndeployFailure        LCMState = iota
	PrologMigratePoweroff        LCMState = iota
	PrologMigratePoweroffFailure LCMState = iota
	PrologMigrateSuspend         LCMState = iota
	PrologMigrateSuspendFailure  LCMState = iota
	BootUndeployFailure          LCMState = iota
	BootStoppedFailure           LCMState = iota
	PrologResumeFailure          LCMState = iota
	PrologUndeployFailure        LCMState = iota
	DiskSnapshotPoweroff         LCMState = iota
	DiskSnapshotRevertPoweroff   LCMState = iota
	DiskSnapshotDeletePoweroff   LCMState = iota
	DiskSnapshotSuspended        LCMState = iota
	DiskSnapshotRevertSuspended  LCMState = iota
	DiskSnapshotDeleteSuspended  LCMState = iota
	DiskSnapshot                 LCMState = iota
	DiskSnapshotDelete           LCMState = iota
	PrologMigrateUnknown         LCMState = iota
	PrologMigrateUnknownFailure  LCMState = iota
	DiskResize                   LCMState = iota
	DiskResizePoweroff           LCMState = iota
	DiskResizeUndeployed         LCMState = iota
)

var lcmStates = map[string]LCMState{
	"LcmInit":                      LcmInit,
	"Prolog":                       Prolog,
	"Boot":                         Boot,
	"Running":                      Running,
	"Migrate":                      Migrate,
	"SaveStop":                     SaveStop,
	"SaveSuspend":                  SaveSuspend,
	"SaveMigrate":                  SaveMigrate,
	"PrologMigrate":                PrologMigrate,
	"PrologResume":                 PrologResume,
	"EpilogStop":                   EpilogStop,
	"Epilog":                       Epilog,
	"Shutdown":                     Shutdown,
	"CleanupResubmit":              CleanupResubmit,
	"Unknown":                      Unknown,
	"Hotplug":                      Hotplug,
	"ShutdownPoweroff":             ShutdownPoweroff,
	"BootUnknown":                  BootUnknown,
	"BootPoweroff":                 BootPoweroff,
	"BootSuspended":                BootSuspended,
	"BootStopped":                  BootStopped,
	"CleanupDelete":                CleanupDelete,
	"HotplugSnapshot":              HotplugSnapshot,
	"HotplugNic":                   HotplugNic,
	"HotplugSaveas":                HotplugSaveas,
	"HotplugSaveasPoweroff":        HotplugSaveasPoweroff,
	"HotplugSaveasSuspended":       HotplugSaveasSuspended,
	"ShutdownUndeploy":             ShutdownUndeploy,
	"EpilogUndeploy":               EpilogUndeploy,
	"PrologUndeploy":               PrologUndeploy,
	"BootUndeploy":                 BootUndeploy,
	"HotplugPrologPoweroff":        HotplugPrologPoweroff,
	"HotplugEpilogPoweroff":        HotplugEpilogPoweroff,
	"BootMigrate":                  BootMigrate,
	"BootFailure":                  BootFailure,
	"BootMigrateFailure":           BootMigrateFailure,
	"PrologMigrateFailure":         PrologMigrateFailure,
	"PrologFailure":                PrologFailure,
	"EpilogFailure":                EpilogFailure,
	"EpilogStopFailure":            EpilogStopFailure,
	"EpilogUndeployFailure":        EpilogUndeployFailure,
	"PrologMigratePoweroff":        PrologMigratePoweroff,
	"PrologMigratePoweroffFailure": PrologMigratePoweroffFailure,
	"PrologMigrateSuspend":         PrologMigrateSuspend,
	"PrologMigrateSuspendFailure":  PrologMigrateSuspendFailure,
	"BootUndeployFailure":          BootUndeployFailure,
	"BootStoppedFailure":           BootStoppedFailure,
	"PrologResumeFailure":          PrologResumeFailure,
	"PrologUndeployFailure":        PrologUndeployFailure,
	"DiskSnapshotPoweroff":         DiskSnapshotPoweroff,
	"DiskSnapshotRevertPoweroff":   DiskSnapshotRevertPoweroff,
	"DiskSnapshotDeletePoweroff":   DiskSnapshotDeletePoweroff,
	"DiskSnapshotSuspended":        DiskSnapshotSuspended,
	"DiskSnapshotRevertSuspended":  DiskSnapshotRevertSuspended,
	"DiskSnapshotDeleteSuspended":  DiskSnapshotDeleteSuspended,
	"DiskSnapshot":                 DiskSnapshot,
	"DiskSnapshotDelete":           DiskSnapshotDelete,
	"PrologMigrateUnknown":         PrologMigrateUnknown,
	"PrologMigrateUnknownFailure":  PrologMigrateUnknownFailure,
	"DiskResize":                   DiskResize,
	"DiskResizePoweroff":           DiskResizePoweroff,
	"DiskResizeUndeployed":         DiskResizeUndeployed,
}

// GetLCMState returns an LCMState for a given string
func GetLCMState(state string) LCMState {
	return lcmStates[state]
}

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
	RunningVMs  int      `xml:"RUNNING_VMS"`
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
	LCMState     LCMState     `xml:"LCM_STATE"`
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
