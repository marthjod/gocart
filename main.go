package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/marthjod/gocart/hostpool"
	"github.com/marthjod/gocart/vmpool"
)

func main() {
	var (
		err          error
		verbose      bool
		elapsed      time.Duration
		vmPoolFile   string
		hostPoolFile string
		cluster      string
		xmlHostFile  *os.File
		xmlVmFile    *os.File
		vmPool       *vmpool.VmPool
		hostPool     *hostpool.HostPool
	)

	flag.StringVar(&vmPoolFile, "vm-pool", "", `VM pool XML dump file path`)
	flag.StringVar(&hostPoolFile, "host-pool", "", `Host pool XML dump file path`)
	flag.StringVar(&cluster, "cluster", "", "Cluster name for host pool lookups")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.Parse()

	xmlHostFile, err = os.Open(hostPoolFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlHostFile.Close()

	xmlVmFile, err = os.Open(vmPoolFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlVmFile.Close()

	hostdata, err := ioutil.ReadAll(xmlHostFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	vmdata, err := ioutil.ReadAll(xmlVmFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	vmPool = vmpool.NewVmPool()

	if elapsed, err = vmPool.Read(vmdata); err != nil {
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

	hostPool = hostpool.NewHostPool()

	if elapsed, err = hostPool.Read(hostdata); err != nil {
		fmt.Println("Error during unmarshaling:", err)
		return
	}

	fmt.Printf("Read in host pool of length %v in %v\n", len(hostPool.Hosts), elapsed)
	if verbose {
		for i := 0; i < len(hostPool.Hosts); i++ {
			host := hostPool.Hosts[i]
			fmt.Printf("%v %v\n", host.Id, host.Template.Datacenter)
		}
	}
	clusterHosts := hostPool.GetHostsInCluster(cluster)
	clusterHosts.MapVms(vmPool)

	for _, h := range clusterHosts.Hosts {
		fmt.Printf("Host %q has VMs\n", h.Name)
		for _, vm := range h.VmPool.Vms {
			fmt.Printf("%s\n", vm.Name)
		}
		fmt.Printf("# of vms: %d\n", len(h.VmPool.Vms))
	}

	billingVms, err := vmPool.GetVmByName("^bil_.+")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("showing all billing vms")
	for _, bvm := range billingVms.Vms {
		fmt.Println(bvm.Name)
	}

}
