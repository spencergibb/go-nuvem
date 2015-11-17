package noop

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/builder"
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
		fmt.Errorf("StaticServerList already inited: %s", s.Namespace)
		return
	}
	s.Namespace = namespace
}

func (s *NoopLoadBalancer) Choose() *loadbalancer.Server {
	return nil
}

var NoopLoadBalancerKey = "NoopLoadBalancer"

func Load() {}

func init() {
	builder.Register(NoopLoadBalancerKey, NewNoopLoadBalancer)
}
