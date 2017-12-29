package dspool

import (
	"encoding/xml"
	"fmt"

	"github.com/marthjod/gocart/ocatypes"
)

// DSPool is a list of Virtual Networks
type DSPool struct {
	XMLName    xml.Name              `xml:"DATASTORE_POOL"`
	Datastores []*ocatypes.Datastore `xml:"DATASTORE"`
}

// ExistsName does this datastore exists?
func (vt *DSPool) ExistsName(n string) bool {
	for _, ds := range vt.Datastores {
		if ds.Name == n {
			return true
		}
	}
	return false
}

// ExistsID does this datastore exists?
func (vt *DSPool) ExistsID(n int) bool {
	for _, ds := range vt.Datastores {
		if ds.ID == n {
			return true
		}
	}
	return false
}

// GetClusterIDByName returns Cluster ID
func (vt *DSPool) GetClusterIDByName(n string) (c int, err error) {
	for _, ds := range vt.Datastores {
		if ds.Name == n {
			return ds.ClusterID, nil
		}
	}
	return -1, fmt.Errorf("could not find datastore with name %s", n)
}

// GetClustersByName returns list of clusters and error
// in opennebula 5 multiple clusters are usable
//func (vt *DSPool) GetClustersByName(n string) (c []int, err error) {
//	var a []int
//	for _, ds := range vt.Datastores {
//		if ds.Name == n {
//			for _, cl := range ds.Clusters {
//				a = append(a, cl.ID)
//			}
//		}
//	}
//	if len(c) > 0 {
//		return a, nil
//	}
//	return a, fmt.Errorf("datastore %s not found", n)
//}

// ApiMethod implements the api.Endpointer interface
func (vt *DSPool) ApiMethod() string {
	return "one.datastorepool.info"
}

// Unmarshal implements the api.Endpointer interface
func (vt *DSPool) Unmarshal(data []byte) error {
	err := xml.Unmarshal(data, vt)
	return err
}

// ApiArgs implements the api.Endpointer interface
// API parameter documentation: http://docs.opennebula.org/4.10/integration/system_interfaces/api.html#one-template-info
func (vt *DSPool) ApiArgs(authstring string) []interface{} {
	return []interface{}{authstring, -2, -1, -1}
}

// NewDSPool creates a new VMTemplatePool
func NewDSPool() *DSPool {
	p := new(DSPool)
	return p
}
