package noop

import (
	"github.com/spencergibb/go-nuvem/loadbalancer"
)

func NewBuilder(ns string) Builder {
	b := Builder{namespace: ns}
	return b
}

type Builder struct {
	namespace string
	Servers   []string
}

func (b Builder) Build() loadbalancer.LoadBalancer {
	lb := NoopLoadBalancer{}
	lb.Namespace = b.namespace
	return &lb
}
