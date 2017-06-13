package ocatypes_test

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/marthjod/gocart/ocatypes"
	"os"
	"github.com/marthjod/gocart/hostpool"
)

var tags = ocatypes.Tags{
	ocatypes.Tag{
		XMLName: xml.Name{"", "FOO"},
		Content: "FOO_VALUE",
	},
}

var xmlInputs = []struct {
	tags      ocatypes.Tags
	lookupKey string
	content   string
	err       error
}{
	{tags, "FOO", "FOO_VALUE", nil}, // successful lookup
	{tags, "BAR", "", errors.New("tag BAR not found")},
}

func TestGetCustom(t *testing.T) {
	for _, expected := range xmlInputs {
		content, err := expected.tags.GetCustom(expected.lookupKey)

		if err == nil && expected.err != nil {
			t.Fatalf("Call did not throw an error")
		}
		if err != nil && expected.err != nil && err.Error() != expected.err.Error() {
			t.Fatalf("Errors do not match: %s != %s", err.Error(), expected.err.Error())
		}
		if content != expected.content {
			t.Fatalf("Contents do not match: %s != %s", content, expected.content)
		}
	}
}

func TestHostIsEmpty(t *testing.T) {
	fixture := "testdata/hostpool.xml"
	f, err := os.Open(fixture)
	defer f.Close()
	if err != nil {
		t.Fatal(err.Error())
	}

	pool, err := hostpool.FromReader(f)
	if err != nil {
		t.Fatal(err.Error())
	}

	emptyHostFound := false
	for _, host := range pool.Hosts {
		if len(host.VmIds) == 0 {
			emptyHostFound = true
			if !host.IsEmpty() {
				t.Error("IsEmpty() should return true for host without VMs")
			}
		} else {
			if host.IsEmpty() {
				t.Error("IsEmpty() should return false for host with VMs")
			}
		}
	}

	if !emptyHostFound {
		t.Errorf("No empty host found in %q for testing IsEmpty()", fixture)
	}
}
