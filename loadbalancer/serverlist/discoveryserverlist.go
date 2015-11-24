package serverlist

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/discovery"
	"github.com/spencergibb/go-nuvem/loadbalancer"
)

type DiscoveryServerList struct {
	Namespace string
	Discovery discovery.Discovery
}

func (s *DiscoveryServerList) Configure(namespace string) {
	if s.Namespace != "" {
		//TODO: use logging
		fmt.Errorf("%s already inited: %s", DisoveryFactoryKey, s.Namespace)
		return
	}
	s.Namespace = namespace
	s.Discovery = discovery.Create(namespace)
}

func (s *DiscoveryServerList) GetServers() []loadbalancer.Server {
	instances := s.Discovery.GetIntances()
	servers := make([]loadbalancer.Server, len(instances))

	for i, instance := range instances {
		servers[i] = loadbalancer.Server{Host: instance.Host, Port: instance.Port}
	}

	return servers
}

func NewDiscoveryServerList() ServerList {
	return &DiscoveryServerList{}
}

var DisoveryFactoryKey = "DiscoveryServerList"

func init() {
	println("registering discovery serverlist")
	Register(DisoveryFactoryKey, NewDiscoveryServerList)
}
