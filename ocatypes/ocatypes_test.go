package ocatypes_test

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/marthjod/gocart/ocatypes"
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
