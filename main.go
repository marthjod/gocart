package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"

	"github.com/marthjod/gocart/api"
	"github.com/marthjod/gocart/hostpool"
	"github.com/marthjod/gocart/vmpool"
)

func main() {
	var (
		verbose    bool
		cluster    string
		cpuprofile string
		user       string
		password   string
		url        string
	)

	flag.StringVar(&cluster, "cluster", "", "Cluster name for host pool lookups")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")
	flag.StringVar(&user, "user", "", `OpenNebula User`)
	flag.StringVar(&password, "password", "", `OpenNebula Password`)
	flag.StringVar(&url, "url", "https://localhost:61443/RPC2", "OpenNebula XML-RPC API URL")

	flag.Parse()

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	apiClient, err := api.NewClient(url, user, password, tr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	vmPool := vmpool.NewVmPool()
	if err := apiClient.Call(vmPool); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	hostPool := hostpool.NewHostPool()
	if err := apiClient.Call(hostPool); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if verbose {
		for i := 0; i < len(vmPool.Vms); i++ {
			vm := vmPool.Vms[i]
			fmt.Printf("%v %v (CPU: %v, template/mem: %v)\n",
				vm.Id, vm.Name, vm.Cpu, vm.Template.Memory)
		}
	}

	hostPool.MapVms(vmPool)

	if verbose {
		for i := 0; i < len(hostPool.Hosts); i++ {
			host := hostPool.Hosts[i]
			fmt.Printf("%v %v\n", host.Id, host.Template.Datacenter)
		}
	}
	clusterHosts := hostPool.GetHostsInCluster(cluster)

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
		acsFQDN, err := bvm.UserTemplate.Items.GetCustom("ACS_FQDN")
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Printf("%s\n", acsFQDN)

	}

}
