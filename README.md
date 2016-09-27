# gocart

## Install

```bash
mkdir -p $GOPATH/src/ocatypes
go build ocatypes.go
cp ocatypes.go $GOPATH/src/ocatypes
go install ocatypes
```

## Run

```bash
go run main.go
  -v	Verbose mode
  -vm-pool string
    	VM pool XML dump file path

go run main.go -vm-pool vmpool.xml
Read in VM pool of length 3 in 956.976µs

go run main.go -vm-pool vmpool.xml -v
Read in VM pool of length 3 in 957.599µs
6 vm-example (CPU: 1, template/mem: 128)
7 vm-in (CPU: 1, template/mem: 128)
8 vm-in (CPU: 1, template/mem: 128)
```

(_vmpool.xml_ from [python-oca](https://github.com/python-oca/python-oca/blob/master/oca/tests/fixtures/vmpool.xml).)

