package builder

import (
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/util"
)

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Build(namespace string) loadbalancer.LoadBalancer {
	result := factories.CallFactory("loadbalancer", namespace, "NoopLoadBalancer")

	lb := result.Interface().(loadbalancer.LoadBalancer)
	lb.Configure(namespace)

	return lb
}
