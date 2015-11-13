package loadbalancer

import (
	"github.com/spencergibb/go-nuvem/util"
)

type (
	LoadBalancer interface {
		util.Initable
		Choose() Server
	}

	// see http://www.captaincodeman.com/2015/03/05/dependency-injection-in-go-golang/
	// for explanation of DI in go
	loadBalancerFactory func() LoadBalancer
)

var (
	New loadBalancerFactory
)