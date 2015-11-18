package simple

import (
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/rule"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
)

func NewBuilder() *Builder {
	b := Builder{}
	return &b
}

type Builder struct {
	namespace  string
	serverList serverlist.ServerList
	rule       rule.Rule
}

func (b *Builder) Namespace(ns string) *Builder {
	b.namespace = ns
	return b
}

func (b *Builder) ServerList(serverList serverlist.ServerList) *Builder {
	b.serverList = serverList
	return b
}

func (b *Builder) Rule(rule rule.Rule) *Builder {
	b.rule = rule
	return b
}

func (b *Builder) Build() loadbalancer.LoadBalancer {
	lb := SimpleLoadBalancer{}
	//TODO: validation
	lb.Namespace = b.namespace
	lb.ServerList = b.serverList
	lb.Rule = b.rule
	return &lb
}
