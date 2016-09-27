// main.go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"ocatypes"
	"os"
	"time"
)

func main() {
	var (
		err        error
		verbose    bool
		elapsed    time.Duration
		vmPoolFile string
		xmlFile    *os.File
		vmPool     *ocatypes.VmPool
	)

	flag.StringVar(&vmPoolFile, "vm-pool", "", `VM pool XML dump file path`)
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.Parse()

	if vmPoolFile == "" {
		flag.PrintDefaults()
		return
	}

	xmlFile, err = os.Open(vmPoolFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()

	data, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	vmPool = ocatypes.NewVmPool()

	if elapsed, err = vmPool.Read(data); err != nil {
		fmt.Println("Error during unmarshaling:", err)
		return
	}

	fmt.Printf("Read in VM pool of length %v in %v\n", len(vmPool.Vms), elapsed)
	if verbose {
		for i := 0; i < len(vmPool.Vms); i++ {
			vm := vmPool.Vms[i]
			fmt.Printf("%v %v (CPU: %v, template/mem: %v)\n",
				vm.Id, vm.Name, vm.Cpu, vm.Template.Memory)
		}
	}
}
