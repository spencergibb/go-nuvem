package noop

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/factory"
)

type (
	NoopLoadBalancer struct {
		Namespace string
	}
)

func NewNoopLoadBalancer() loadbalancer.LoadBalancer {
	return &NoopLoadBalancer{}
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

func Load() {}

func init() {
	factory.Register(FactoryKey, NewNoopLoadBalancer)
}
