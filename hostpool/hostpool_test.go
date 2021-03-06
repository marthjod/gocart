package hostpool_test

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/marthjod/gocart/host"
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

func getHostPoolFromFile(path string) *hostpool.HostPool {
	pool := &hostpool.HostPool{}
	c := getFileContents(path)
	_ = xml.Unmarshal(c, &pool)
	return pool
}

func TestFromReader(t *testing.T) {
	var expected = 2

	f, err := os.Open("testdata/hostpool.xml")
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		t.Fatal(err.Error())
	}

	pool, err := hostpool.FromReader(f)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(pool.Hosts) != expected {
		t.Errorf("Host pool has length %d, expected %d", len(pool.Hosts), expected)
	}

}

func TestGetHostsInCluster(t *testing.T) {
	pool := getHostPoolFromFile("testdata/hostpool.xml")

	for _, expected := range expectedPools {
		clusterPool := pool.GetHostsInCluster(expected.cluster)
		if len(clusterPool.Hosts) != expected.poolLen {
			t.Fatalf("Unexpected pool size (%d != %d) for cluster %s",
				len(clusterPool.Hosts), expected.poolLen, expected.cluster)
		}
	}
}

func TestFilterHostsByStates(t *testing.T) {
	pool := getHostPoolFromFile("testdata/hostpool.xml")

	disabledHosts := pool.FilterHostsByStates(host.Disabled)
	if len(disabledHosts.Hosts) != 1 {
		t.Fatalf("Expected 1 disabled host, found %d", len(disabledHosts.Hosts))
	}
	if disabledHosts.Hosts[0].Name != "thost" {
		t.Fatalf("Found wrong disabled host %s", disabledHosts.Hosts[0].Name)
	}

	emptyPool := pool.FilterHostsByStates(host.Error)
	if len(emptyPool.Hosts) > 0 {
		t.Fatalf("Found more than 0 hosts in state ERROR")
	}

	twoStates := pool.FilterHostsByStates(host.Disabled, host.Monitored)
	if len(twoStates.Hosts) != 2 {
		t.Fatalf("Expected 2 hosts for states Disabled + ERROR, found %d", len(twoStates.Hosts))
	}
}

func TestFilterOutEmptyHosts(t *testing.T) {
	pool := getHostPoolFromFile("testdata/hostpool.xml")
	hostsRunningVMs := pool.FilterOutEmptyHosts()
	if len(hostsRunningVMs.Hosts) > 1 {
		t.Error("Found too many non-empty hosts")
	}
}

func TestFilterChain(t *testing.T) {
	pool := getHostPoolFromFile("testdata/hostpool.xml")
	disabledHosts := pool.FilterHostsByStates(host.Disabled)
	if len(disabledHosts.Hosts) != 1 {
		t.Errorf("Expected 1 disabled host, found %d", len(disabledHosts.Hosts))
	}

	filtered := pool.FilterHostsByStates(host.Disabled).FilterOutEmptyHosts()
	if len(filtered.Hosts) > 1 {
		t.Error("Found too many hosts while chaining filters")
	}
}
