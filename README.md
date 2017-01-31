# gocart

## Install

```
go get -u github.com/marthjod/gocart
```

## Run tests

```
go test -cover ./...
```

## Run

```
go run main.go -h
  -cluster string
    	Cluster name for host pool lookups
  -cpuprofile string
    	write cpu profile to file
  -password string
    	OpenNebula Password
  -url string
    	OpenNebula XML-RPC API URL (default "https://localhost:61443/RPC2")
  -user string
    	OpenNebula User
  -v	Verbose mode
```

