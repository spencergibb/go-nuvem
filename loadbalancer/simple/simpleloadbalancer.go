package simple

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/factory"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
	"math/rand"
)

type (
	SimpleLoadBalancer struct {
		Namespace  string
		ServerList serverlist.ServerList
	}
)

func NewSimpleLoadBalancer() loadbalancer.LoadBalancer {
	return &SimpleLoadBalancer{}
}

func (s *SimpleLoadBalancer) Configure(namespace string) {
	if s.Namespace != "" {
		//TODO: use logging
		fmt.Errorf("%s already inited: %s", FactoryKey, s.Namespace)
		return
	}
	s.Namespace = namespace
}

func (s *SimpleLoadBalancer) Choose() *loadbalancer.Server {
	servers := s.ServerList.GetServers()
	//	TODO: implement rules
	idx := rand.Intn(len(servers))
	return &servers[idx]
}

var FactoryKey = "SimpleLoadBalancer"

func Load() {}

func init() {
	factory.Register(FactoryKey, NewSimpleLoadBalancer)
}
