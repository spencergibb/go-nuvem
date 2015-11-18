package noop

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
)

type NoopLoadBalancer struct {
	Namespace string
}

func (s *NoopLoadBalancer) Configure(namespace string) {
	if s.Namespace != "" {
		//TODO: use logging
		fmt.Errorf("%s already inited: %s", FactoryKey, s.Namespace)
		return
	}
	s.Namespace = namespace
}

func (s *NoopLoadBalancer) Choose() *loadbalancer.Server {
	return nil
}

var FactoryKey = "NoopLoadBalancer"

func NewNoopLoadBalancer() loadbalancer.LoadBalancer {
	return &NoopLoadBalancer{}
}

func init() {
	loadbalancer.Register(FactoryKey, NewNoopLoadBalancer)
}
