package clusterpool

import (
	"encoding/xml"
	"fmt"

	"github.com/marthjod/gocart/ocatypes"
)

// ImagePool is a list of VMTemplates
type ClusterPool struct {
	XMLName  xml.Name            `xml:"CLUSTER_POOL"`
	Clusters []*ocatypes.Cluster `xml:"CLUSTER"`
}

// ApiMethod implements the api.Endpointer interface
func (vt *ClusterPool) ApiMethod() string {
	return "one.clusterpool.info"
}

// Unmarshal implements the api.Endpointer interface
func (vt *ClusterPool) Unmarshal(data []byte) error {
	err := xml.Unmarshal(data, vt)
	return err
}

// ApiArgs implements the api.Endpointer interface
// API parameter documentation: http://docs.opennebula.org/4.10/integration/system_interfaces/api.html
func (vt *ClusterPool) ApiArgs(authstring string) []interface{} {
	return []interface{}{authstring, -2, -1, -1}
}

// NewClusterPool creates a new ClusterPool
func NewClusterPool() *ClusterPool {
	p := new(ClusterPool)
	return p
}

// ExistsName does this cluster exists?
func (vt *ClusterPool) ExistsName(n string) bool {
	for _, cl := range vt.Clusters {
		if cl.Name == n {
			return true
		}
	}
	return false
}

// ExistsID does this cluster exists?
func (vt *ClusterPool) ExistsID(n int) bool {
	for _, cl := range vt.Clusters {
		if cl.ID == n {
			return true
		}
	}
	return false
}

// GetIDByName returns ID of Cluster or -1 and error
func (vt *ClusterPool) GetIDByName(name string) (i int, err error) {
	for _, cl := range vt.Clusters {
		if cl.Name == name {
			return cl.ID, nil
		}
	}
	return -1, fmt.Errorf("could not find cluster with name %s", name)
}
