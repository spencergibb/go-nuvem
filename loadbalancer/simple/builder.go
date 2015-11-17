package simple

import (
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
)

func NewBuilder() *Builder {
	b := Builder{}
	return &b
}

type Builder struct {
	namespace  string
	serverList serverlist.ServerList
}

func (b *Builder) Namespace(ns string) *Builder {
	b.namespace = ns
	return b
}

func (b *Builder) ServerList(serverList serverlist.ServerList) *Builder {
	b.serverList = serverList
	return b
}

func (b *Builder) Build() loadbalancer.LoadBalancer {
	lb := SimpleLoadBalancer{}
	lb.Namespace = b.namespace
	lb.ServerList = b.serverList
	return &lb
}
