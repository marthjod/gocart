package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bogue1979/xmlrpc"
)

type XMLRPCEndpointer interface {
	ApiMethod() string
	ApiArgs(string) []interface{}
	Unmarshal([]byte) error
}

type Rpc struct {
	Client     *xmlrpc.Client
	Url        string
	AuthString string
}

// NewClient creates a RPC connection
func NewClient(url, user, password string, transport http.RoundTripper, timeout time.Duration) (*Rpc, error) {
	client, err := xmlrpc.NewClient(url, transport, timeout)
	if err != nil {
		return nil, err
	}
	return newClient(url, user, password, client)
}

func newClient(url, user, password string, client *xmlrpc.Client) (*Rpc, error) {
	return &Rpc{
		Client:     client,
		Url:        url,
		AuthString: fmt.Sprintf("%s:%s", user, password),
	}, nil
}

// Call the endpoint
func (c *Rpc) Call(endpoint XMLRPCEndpointer) error {
	args := endpoint.ApiArgs(c.AuthString)
	method := endpoint.ApiMethod()
	result := []interface{}{}

	if err := c.Client.Call(method, args, &result); err != nil {
		return err
	}

	apiCallSucceeded, ok := result[0].(bool)
	if !ok {
		return fmt.Errorf("malformed XMLRPC response")
	}
	if !apiCallSucceeded {
		// panic: interface conversion: interface is bool, not string
		return fmt.Errorf("API call against %s unsuccessful", c.Url)
	}
	if w, ok := result[1].(string); ok {
		endpoint.Unmarshal([]byte(w))
		return nil
	}
	return fmt.Errorf("no know result type received from RPC call")
}
