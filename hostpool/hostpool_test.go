package hostpool_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/marthjod/gocart/hostpool"
)

var expectedPools = []struct {
	cluster string
	poolLen int
}{
	{"default", 2},
	{"non-existent", 0},
}

func getFileContents(path string) []byte {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}
	}
	return f
}

func TestApiMethod(t *testing.T) {
	const apiMethod = "one.hostpool.info"

	if hostpool.NewHostPool().ApiMethod() != apiMethod {
		t.Fatalf("API method differs from %s", apiMethod)
	}
}

func TestApiArgs(t *testing.T) {
	const (
		authstring = "user:pass"
		expected   = "[user:pass]"
	)

	args := hostpool.NewHostPool().ApiArgs(authstring)
	argsStr := fmt.Sprintf("%s", args)
	if argsStr != expected {
		t.Fatalf("Mismatch: %s != %s", argsStr, expected)
	}

}

func TestUnmarshalValidHostpool(t *testing.T) {
	pool := hostpool.NewHostPool()
	xml := getFileContents("testdata/hostpool.xml")

	err := pool.Unmarshal(xml)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestUnmarshalInvalidXML(t *testing.T) {
	var expected = "expected element type <HOST_POOL> but have <INVALID_HOST_POOL>"
	pool := hostpool.NewHostPool()
	xml := getFileContents("testdata/invalid-hostpool.xml")

	err := pool.Unmarshal(xml)
	if err == nil {
		t.Fatal("Call did not throw an error")
	} else {
		if err.Error() != expected {
			t.Fatalf("Errors do not match: %s != %s", err.Error(), expected)
		}
	}
}

func TestGetHostsInCluster(t *testing.T) {
	pool := hostpool.NewHostPool()
	xml := getFileContents("testdata/hostpool.xml")
	pool.Unmarshal(xml)

	for _, expected := range expectedPools {
		clusterPool := pool.GetHostsInCluster(expected.cluster)
		if len(clusterPool.Hosts) != expected.poolLen {
			t.Fatalf("Unexpected pool size (%d != %d) for cluster %s",
				len(clusterPool.Hosts), expected.poolLen, expected.cluster)
		}
	}

}
