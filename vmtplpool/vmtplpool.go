package vmtplpool

import (
	"encoding/xml"

	"github.com/marthjod/gocart/ocatypes"
)

// VMTemplatePool is a list of VMTemplates
type VMTemplatePool struct {
	XMLName   xml.Name               `xml:"VMTEMPLATE_POOL"`
	Templates []*ocatypes.VmTemplate `xml:"VMTEMPLATE"`
}

// ApiMethod implements the api.Endpointer interface
func (vt *VMTemplatePool) ApiMethod() string {
	return "one.templatepool.info"
}

// Unmarshal implements the api.Endpointer interface
func (vt *VMTemplatePool) Unmarshal(data []byte) error {
	err := xml.Unmarshal(data, vt)
	return err
}

// ApiArgs implements the api.Endpointer interface
// API parameter documentation: http://docs.opennebula.org/4.10/integration/system_interfaces/api.html#one-template-info
func (vt *VMTemplatePool) ApiArgs(authstring string) []interface{} {
	return []interface{}{authstring, -2, -1, -1}
}

// NewVMTemplatePool creates a new VMTemplatePool
func NewVMTemplatePool() *VMTemplatePool {
	p := new(VMTemplatePool)
	return p
}
