package imagepool

import (
	"encoding/xml"

	"github.com/marthjod/gocart/ocatypes"
)

// ImagePool is a list of VMTemplates
type ImagePool struct {
	XMLName xml.Name          `xml:"IMAGE_POOL"`
	Images  []*ocatypes.Image `xml:"IMAGE"`
}

// ApiMethod implements the api.Endpointer interface
func (vt *ImagePool) ApiMethod() string {
	return "one.imagepool.info"
}

// Unmarshal implements the api.Endpointer interface
func (vt *ImagePool) Unmarshal(data []byte) error {
	err := xml.Unmarshal(data, vt)
	return err
}

// ApiArgs implements the api.Endpointer interface
// API parameter documentation: http://docs.opennebula.org/4.10/integration/system_interfaces/api.html#one-template-info
func (vt *ImagePool) ApiArgs(authstring string) []interface{} {
	return []interface{}{authstring, -2, -1, -1}
}

// NewImagePool creates a new ImagePool
func NewImagePool() *ImagePool {
	p := new(ImagePool)
	return p
}

// ExistsName does this image exists?
func (vt *ImagePool) ExistsName(n string) bool {
	for _, image := range vt.Images {
		if image.Name == n {
			return true
		}
	}
	return false
}

// ExistsID does this image exists?
func (vt *ImagePool) ExistsID(i int) bool {
	for _, image := range vt.Images {
		if image.ID == i {
			return true
		}
	}
	return false
}
