package vm

import (
	"encoding/xml"
	"io"

	"github.com/marthjod/gocart/ocatypes"
)

// FromReader reads into a VM struct.
func FromReader(r io.Reader) (*ocatypes.VM, error) {
	v := ocatypes.VM{}
	dec := xml.NewDecoder(r)
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}
