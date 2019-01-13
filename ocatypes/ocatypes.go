package ocatypes

import (
	"encoding/xml"
	"fmt"
)

// State represents VM state
type State int

// LCMState represents LCM (lifecycle manager) state
type LCMState int

//go:generate stringer -type=State
const (
	Init       State = iota
	Pending    State = iota
	Hold       State = iota
	Active     State = iota
	Stopped    State = iota
	Suspended  State = iota
	Done       State = iota
	Failed     State = iota
	Poweroff   State = iota
	Undeployed State = iota
)

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

// States maps VM state names to their constant State values.
var States = map[string]State{
	"Init":       Init,
	"Pending":    Pending,
	"Hold":       Hold,
	"Active":     Active,
	"Stopped":    Stopped,
	"Suspended":  Suspended,
	"Done":       Done,
	"Failed":     Failed,
	"Poweroff":   Poweroff,
	"Undeployed": Undeployed,
}

// LCMStates maps LCM state names to their constant LCMState values.
var LCMStates = map[string]LCMState{
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

// GetState returns a VM state for a given string.
func GetState(state string) State {
	return States[state]
}

// GetLCMState returns an LCMState for a given string.
func GetLCMState(state string) LCMState {
	return LCMStates[state]
}

// ClusterPool is a list of clusters.
type ClusterPool struct {
	XMLName  xml.Name   `xml:"CLUSTER_POOL"`
	Clusters []*Cluster `xml:"CLUSTER"`
}

// Cluster represents a cluster.
type Cluster struct {
	XMLName      xml.Name `xml:"CLUSTER"`
	Name         string   `xml:"NAME"`
	ID           int      `xml:"ID"`
	DatastoreIDs []int    `xml:"DATASTORES"`
	VnetIDs      []int    `xml:"VNETS"`
}

// DSPool is a list of Datastores.
type DSPool struct {
	XMLName    xml.Name     `xml:"DATASTORE_POOL"`
	Datastores []*Datastore `xml:"DATASTORE"`
}

// Datastore represents a datastore.
type Datastore struct {
	XMLName   xml.Name `xml:"DATASTORE"`
	Name      string   `xml:"NAME"`
	ID        int      `xml:"ID"`
	ClusterID int      `xml:"CLUSTER_ID"`
	Cluster   string   `xml:"CLUSTER"`
}

// VNetPool is a list of Virtual Networks.
type VNetPool struct {
	XMLName  xml.Name `xml:"VNET_POOL"`
	Networks []*VNet  `xml:"VNET"`
}

// VNet represents a virtual network.
type VNet struct {
	XMLName   xml.Name `xml:"VNET"`
	Name      string   `xml:"NAME"`
	ID        int      `xml:"ID"`
	Cluster   string   `xml:"CLUSTER"`
	ClusterID int      `xml:"CLUSTER_ID"`
	Bridge    string   `xml:"BRIDGE"`
}

// Disk represents a disk.
type Disk struct {
	XMLName xml.Name `xml:"DISK"`
	Name    string   `xml:"IMAGE"`
	ID      int      `xml:"IMAGE_ID"`
}

// Image represents an image.
type Image struct {
	XMLName     xml.Name `xml:"IMAGE"`
	ID          int      `xml:"ID"`
	Name        string   `xml:"NAME"`
	Datastore   string   `xml:"DATASTORE"`
	DatastoreID int      `xml:"DATASTORE_ID"`
	RunningVMs  int      `xml:"RUNNING_VMS"`
}

// NIC represents a network interface.
type NIC struct {
	XMLName   xml.Name `xml:"NIC"`
	Name      string   `xml:"NETWORK"`
	NetworkID int      `xml:"NETWORK_ID"`
}

// VMTemplatePool is a list of VMTemplates.
type VMTemplatePool struct {
	XMLName   xml.Name      `xml:"VMTEMPLATE_POOL"`
	Templates []*VMTemplate `xml:"VMTEMPLATE"`
}

// HostTemplate represents a host template.
type HostTemplate struct {
	XMLName        xml.Name `xml:"TEMPLATE"`
	CPU            string   `xml:"CPU"`
	Disk           []Disk   `xml:"DISK"`
	Memory         string   `xml:"MEMORY"`
	Name           string   `xml:"NAME"`
	Nics           []NIC    `xml:"NIC"`
	VCPU           string   `xml:"VCPU"`
	Datacenter     string   `xml:"DATACENTER"`
	Requirements   string   `xml:"REQUIREMENTS"`
	DSRequirements string   `xml:"SCHED_DS_REQUIREMENTS"`
	Items          Tags     `xml:",any"`
}

// VMTemplate represents a VM template.
type VMTemplate struct {
	ID       int          `xml:"ID"`
	Name     string       `xml:"NAME"`
	Uname    string       `xml:"UNAME"`
	RegTime  int          `xml:"REGTIME"`
	Template HostTemplate `xml:"TEMPLATE"`
	Memory   int          `xml:"MEMORY"`
	VMID     int          `xml:"VMID"`
	Disk     []Disk       `xml:"DISK"`
	CPU      string       `xml:"CPU"`
}

// VM ...
// Node (current OpenNebula node the VM is running on) determination is best effort:
// in case of multiple hosts, this *might* not be 100% reliable, ie. pick the current one,
// although tests reproducibly showed correct results.
type VM struct {
	XMLName      xml.Name     `xml:"VM"`
	ID           int          `xml:"ID"`
	Name         string       `xml:"NAME"`
	CPU          int          `xml:"CPU"`
	LastPoll     int          `xml:"LAST_POLL"`
	State        State        `xml:"STATE"`
	LCMState     LCMState     `xml:"LCM_STATE"`
	Resched      int          `xml:"RESCHED"`
	DeployID     string       `xml:"DEPLOY_ID"`
	Template     VMTemplate   `xml:"TEMPLATE"`
	UserTemplate UserTemplate `xml:"USER_TEMPLATE"`
	Node         string       `xml:"HISTORY_RECORDS>HISTORY>HOSTNAME"`
}

// UserTemplate represents a user template.
type UserTemplate struct {
	Items Tags `xml:",any"`
}

// Tag is an XML tag.
type Tag struct {
	XMLName xml.Name
	Content string `xml:",chardata"`
}

// Tags is a list of Tags.
type Tags []Tag

// GetCustom returns values from custom-defined XML tags.
func (tags Tags) GetCustom(tagName string) (string, error) {
	for _, tag := range tags {
		if tagName == tag.XMLName.Local {
			return tag.Content, nil
		}
	}
	return "", fmt.Errorf("tag %s not found", tagName)
}

// Host represents an OpenNebula host.
type Host struct {
	XMLName   xml.Name     `xml:"HOST"`
	ID        int          `xml:"ID"`
	Name      string       `xml:"NAME"`
	State     int          `xml:"STATE"`
	Cluster   string       `xml:"CLUSTER"`
	ClusterID int          `xml:"CLUSTER_ID"`
	Template  HostTemplate `xml:"TEMPLATE"`
	VMIDs     []int        `xml:"VMS>ID"`
}

// IsEmpty checks if a host has no VMs.
func (h *Host) IsEmpty() bool {
	return len(h.VMIDs) == 0
}

// VMPool represents a VM pool.
type VMPool struct {
	XMLName xml.Name `xml:"VM_POOL"`
	VMs     []VM     `xml:"VM"` // ?
}
