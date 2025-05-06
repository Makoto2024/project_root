package main

import (
	"errors"
	"flag"
)

type serviceType string

const (
	serviceTypeHelloService serviceType = "helloService"
	serviceTypePongService  serviceType = "pongService"
)

var allServiceTypes = []serviceType{
	serviceTypeHelloService,
	serviceTypePongService,
}

var _ flag.Value = (*serviceType)(nil)

func (t serviceType) String() string {
	return string(t)
}

func (t *serviceType) Get() any {
	return *t
}

func (t *serviceType) Set(val string) error {
	for _, k := range allServiceTypes {
		if string(k) == val {
			*t = k
			return nil
		}
	}
	return errors.New("invalid service type")
}

// cfg will be set via flags.
type cfg struct {
	IP      string
	Port    int
	Service serviceType
	Param   string
}

func (c *cfg) registerFlags(fs *flag.FlagSet) {
	if fs == nil {
		fs = flag.CommandLine
	}
	fs.StringVar(&c.IP, "ip", "127.0.0.1", "IP address to connect to")
	fs.IntVar(&c.Port, "port", 8080, "Port to connect to")
	fs.Var(&c.Service, "service", "Service type")
	fs.StringVar(&c.Param, "param", "DefaultStr", "Parameter")
}
