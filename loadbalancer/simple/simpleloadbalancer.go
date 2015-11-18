package simple

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/factory"
	"github.com/spencergibb/go-nuvem/loadbalancer/rule"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
	slfactory "github.com/spencergibb/go-nuvem/loadbalancer/serverlist/factory"
)

type (
	SimpleLoadBalancer struct {
		Namespace  string
		ServerList serverlist.ServerList
		Rule       rule.Rule
	}
)

func (s *SimpleLoadBalancer) Configure(namespace string) {
	if s.Namespace != "" {
		//TODO: use logging
		fmt.Errorf("%s already inited: %s", FactoryKey, s.Namespace)
		return
	}
	s.ServerList = slfactory.Create(namespace)
	s.Rule = rule.CreateRule(namespace)
	s.Namespace = namespace
}

func (s *SimpleLoadBalancer) Choose() *loadbalancer.Server {
	servers := s.ServerList.GetServers()
	return s.Rule.Choose(servers)
}

var FactoryKey = "SimpleLoadBalancer"

func NewSimpleLoadBalancer() loadbalancer.LoadBalancer {
	return &SimpleLoadBalancer{}
}

func init() {
	factory.Register(FactoryKey, NewSimpleLoadBalancer)
}
