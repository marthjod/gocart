# gocart

[![Travis CI Build Status](https://travis-ci.org/marthjod/gocart.svg?branch=master)](https://travis-ci.org/marthjod/gocart)

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
  -pattern-filter string
        Regexp filter for distinct VM name pattern auto-discovery (default "^([a-z]{2}).+([a-z]{2})$")
  -pattern-filter-infix string
        Infix for distinct VM name patterns (default ".+")
  -pattern-filter-prefix string
        Prefix for distinct VM name patterns (default "^")
  -pattern-filter-suffix string
        Suffix for distinct VM name patterns (default "$")
  -url string
        OpenNebula XML-RPC API URL (default "https://localhost:61443/RPC2")
  -user string
        OpenNebula User
  -v    Verbose mode
```

