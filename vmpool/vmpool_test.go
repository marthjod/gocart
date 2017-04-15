package vmpool_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/marthjod/gocart/vmpool"
	"github.com/marthjod/gocart/ocatypes"
)

func getFileContents(path string) []byte {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}
	}
	return f
}

func getVmPoolFromFile(path string) *vmpool.VmPool {
	pool := vmpool.NewVmPool()
	xml := getFileContents(path)
	_ = pool.Unmarshal(xml)
	return pool
}

func TestApiMethod(t *testing.T) {
	const apiMethod = "one.vmpool.info"

	if vmpool.NewVmPool().ApiMethod() != apiMethod {
		t.Fatalf("API method differs from %s", apiMethod)
	}
}

func TestApiArgs(t *testing.T) {
	const (
		authstring = "user:pass"
		expected   = "[user:pass %!s(int=-2) %!s(int=-1) %!s(int=-1) %!s(int=-1)]"
	)

	args := vmpool.NewVmPool().ApiArgs(authstring)
	argsStr := fmt.Sprintf("%s", args)
	if argsStr != expected {
		t.Fatalf("Mismatch: %s != %s", argsStr, expected)
	}
}

func TestUnmarshalValidVmpool(t *testing.T) {
	pool := vmpool.NewVmPool()
	xml := getFileContents("testdata/vmpool.xml")

	err := pool.Unmarshal(xml)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestUnmarshalInvalidXML(t *testing.T) {
	var expected = "expected element type <VM_POOL> but have <INVALID_VM_POOL>"
	pool := vmpool.NewVmPool()
	xml := getFileContents("testdata/invalid-vmpool.xml")

	err := pool.Unmarshal(xml)
	if err == nil {
		t.Fatal("Call did not throw an error")
	} else {
		if err.Error() != expected {
			t.Fatalf("Errors do not match: %s != %s", err.Error(), expected)
		}
	}
}

func TestFromReader(t *testing.T) {
	var expected = 3

	f, err := os.Open("testdata/vmpool.xml")
	defer f.Close()
	if err != nil {
		t.Fatal(err.Error())
	}

	pool, err := vmpool.FromReader(f)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(pool.Vms) != expected {
		t.Errorf("VM pool has length %d, expected %d", len(pool.Vms), expected)
	}

}

func TestGetVmsById(t *testing.T) {
	var expectedVmsById = map[int]int{
		6:    1,
		7:    1,
		8:    1,
		1337: 0,
	}

	pool := getVmPoolFromFile("testdata/vmpool.xml")

	for id, vmsExpected := range expectedVmsById {
		vmsFound := pool.GetVmsById(id)
		if len(vmsFound.Vms) != vmsExpected {
			t.Errorf("Expected %d VM(s), found %d", vmsExpected, len(vmsFound.Vms))
		}
	}

}

func TestGetVmsByName(t *testing.T) {
	var expectedVmsByPattern = map[string]int{
		"^vm":      3,
		"vm-in$":   2,
		"exa":      1,
		"dummy-vm": 0,
	}

	pool := getVmPoolFromFile("testdata/vmpool.xml")

	for pattern, vmsExpected := range expectedVmsByPattern {
		vmsFound, err := pool.GetVmsByName(pattern)
		if err != nil {
			t.Fatal(err.Error())
		}
		if len(vmsFound.Vms) != vmsExpected {
			t.Errorf("%q: expected %d VM(s), found %d", pattern, vmsExpected, len(vmsFound.Vms))
		}

	}
}

func TestGetDistinctVmNamePatterns(t *testing.T) {
	var prefix = "^"
	var infix = ".+"
	var suffix = "$"
	var expectedDistinctPatterns = []struct {
		filter  string
		pattern string
	}{
		{"^([vm]{2})-([ni]{2})$", "^vm.+in$"},
		{"^vm-(ex)....(e)$", "^ex.+e$"},
		{"^vm-([a-z]{4}).(.).*", "^exam.+l$"},
	}

	pool := getVmPoolFromFile("testdata/vmpool.xml")

	for _, expected := range expectedDistinctPatterns {
		distinctPatterns := pool.GetDistinctVmNamePatterns(expected.filter, prefix, infix, suffix)

		if !distinctPatterns[expected.pattern] {
			t.Errorf("Expected distinct pattern %q not extracted by filter %q", expected.pattern, expected.filter)
		}
	}
}


func TestGetDistinctVmNamePatternsExtractHostname(t *testing.T) {
	var prefix = "^"
	var infix = ".+"
	var suffix = "$"
	var expectedDistinctPatterns = []struct {
		filter  string
		pattern string
	}{
		{".{2}(amp).*(amp)", "^amp.+amp$"},
		{"^..-(..).*(m)$", "^vm.+m$"},
		{"^(.{2}).*(example.com)", "^in.+example.com$"},
	}

	pool := getVmPoolFromFile("testdata/vmpool.xml")
	fqdnExtractor := func(vm *ocatypes.Vm) string {
		h, _ := vm.UserTemplate.Items.GetCustom("CUSTOM_FQDN")
		return h
	}

	for _, expected := range expectedDistinctPatterns {
		distinctPatterns := pool.GetDistinctVmNamePatternsExtractHostname(expected.filter, prefix, infix, suffix, fqdnExtractor)

		if !distinctPatterns[expected.pattern] {
			t.Errorf("Expected distinct pattern %q not extracted by filter %q", expected.pattern, expected.filter)
		}
	}
}
