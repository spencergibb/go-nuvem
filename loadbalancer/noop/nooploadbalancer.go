package noop

import (
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/builder"
)

type (
	NoopLoadBalancer struct {
		namespace string
	}
)

func NewNoopLoadBalancer() loadbalancer.LoadBalancer {
	return &NoopLoadBalancer{}
}

func (s *NoopLoadBalancer) Init(namespace string) {
	s.namespace = namespace
}

func (s *NoopLoadBalancer) Choose() *loadbalancer.Server {
	return nil
}

var NoopLoadBalancerKey = "NoopLoadBalancer"

func Load() {}

func init() {
	builder.Register(NoopLoadBalancerKey, NewNoopLoadBalancer)
}
