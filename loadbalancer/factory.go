package loadbalancer

import (
	"github.com/spencergibb/go-nuvem/util"
)

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Create(namespace string) LoadBalancer {
	result := factories.CallFactory("loadbalancer", namespace, "NoopLoadBalancer")

	lb := result.Interface().(LoadBalancer)
	lb.Configure(namespace)

	return lb
}
