package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bogue1979/xmlrpc"
)

// XMLRPCEndpointer defines behavior of a specific XML-RPC API endpoint.
type XMLRPCEndpointer interface {
	APIMethod() string
	APIArgs(string) []interface{}
	Unmarshal([]byte) error
}

// RPC represents an XML-RPC client.
type RPC struct {
	Client     *xmlrpc.Client
	URL        string
	AuthString string
}

// NewClient returns a prepared XML-RPC client.
func NewClient(url, user, password string, transport http.RoundTripper, timeout time.Duration) (*RPC, error) {
	client, err := xmlrpc.NewClient(url, transport, timeout)
	if err != nil {
		return nil, err
	}
	return newClient(url, user, password, client)
}

func newClient(url, user, password string, client *xmlrpc.Client) (*RPC, error) {
	return &RPC{
		Client:     client,
		URL:        url,
		AuthString: fmt.Sprintf("%s:%s", user, password),
	}, nil
}

// Call issues a request against the API endpoint.
func (c *RPC) Call(endpoint XMLRPCEndpointer) error {
	args := endpoint.APIArgs(c.AuthString)
	method := endpoint.APIMethod()
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
		return fmt.Errorf("API call against %s unsuccessful", c.URL)
	}
	if w, ok := result[1].(string); ok {
		endpoint.Unmarshal([]byte(w))
		return nil
	}
	return fmt.Errorf("no known result type received from RPC call")
}
