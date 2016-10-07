package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/marthjod/gocart/hostpool"
	"github.com/marthjod/gocart/vmpool"
)

func main() {
	var (
		err          error
		verbose      bool
		vmPoolFile   string
		hostPoolFile string
		cluster      string
		xmlHostFile  *os.File
		xmlVmFile    *os.File
		vmPool       *vmpool.VmPool
		hostPool     *hostpool.HostPool
		cpuprofile   string
	)

	flag.StringVar(&vmPoolFile, "vm-pool", "", `VM pool XML dump file path`)
	flag.StringVar(&hostPoolFile, "host-pool", "", `Host pool XML dump file path`)
	flag.StringVar(&cluster, "cluster", "", "Cluster name for host pool lookups")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")

	flag.Parse()

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	xmlHostFile, err = os.Open(hostPoolFile)
	defer xmlHostFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	xmlVmFile, err = os.Open(vmPoolFile)
	defer xmlVmFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	vmPool, err = vmpool.FromReader(xmlVmFile)
	if err != nil {
		panic(err)
	}

	if verbose {
		for i := 0; i < len(vmPool.Vms); i++ {
			vm := vmPool.Vms[i]
			fmt.Printf("%v %v (CPU: %v, template/mem: %v)\n",
				vm.Id, vm.Name, vm.Cpu, vm.Template.Memory)
		}
	}

	hostPool, err = hostpool.FromReader(xmlHostFile)
	if err != nil {
		panic(err)
	}

	hostPool.MapVms(vmPool)

	if verbose {
		for i := 0; i < len(hostPool.Hosts); i++ {
			host := hostPool.Hosts[i]
			fmt.Printf("%v %v\n", host.Id, host.Template.Datacenter)
		}
	}
	clusterHosts := hostPool.GetHostsInCluster(cluster)
	// clusterHosts.MapVms(vmPool)

	for _, h := range clusterHosts.Hosts {
		fmt.Printf("Host %q has VMs\n", h.Name)
		for _, vm := range h.VmPool.Vms {
			fmt.Printf("%s\n", vm.Name)
		}
		fmt.Printf("# of vms: %d\n", len(h.VmPool.Vms))
	}

	billingVms, err := vmPool.GetVmsByName("^bil_.+")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("showing all billing vms")
	for _, bvm := range billingVms.Vms {
		fmt.Println(bvm.Name)
		fmt.Println("User Template:")
		for _, v := range bvm.UserTemplate.Items {
			fmt.Printf("%s = %s\n", v.XMLName.Local, v.Content)
		}

	}
}
