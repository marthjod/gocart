package vm

import (
	"encoding/xml"
	"github.com/marthjod/gocart/ocatypes"
	"io"
)

func FromReader(r io.Reader) (*ocatypes.Vm, error) {
	v := ocatypes.Vm{}
	dec := xml.NewDecoder(r)
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}
