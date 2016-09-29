package ocahostpool

import (
	"encoding/xml"
	"gocart/ocatypes"
	"time"
)

type HostPool struct {
	XMLName xml.Name         `xml:"HOST_POOL"`
	Hosts   []*ocatypes.Host `xml:"HOST"`
}

func NewHostPool() *HostPool {
	p := new(HostPool)
	return p
}

func (hostPool *HostPool) Read(xmlData []byte) (time.Duration, error) {
	var (
		err     error
		elapsed time.Duration
	)

	_, elapsed, err = ocatypes.Read(xmlData, hostPool)
	return elapsed, err
}
