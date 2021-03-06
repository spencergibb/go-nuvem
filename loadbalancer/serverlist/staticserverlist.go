package serverlist

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spf13/viper"
	"net"
	"strconv"
)

type StaticServerList struct {
	Namespace string
	Servers   []string
}

func (s *StaticServerList) Configure(namespace string) {
	if s.Namespace != "" {
		//TODO: use logging
		fmt.Errorf("%s already inited: %s", StaticFactoryKey, s.Namespace)
		return
	}
	s.Namespace = namespace
	serverConfigs := viper.GetStringSlice(s.GetServerKey())
	fmt.Printf("serverConfigs %+v\n", serverConfigs)
	s.Servers = serverConfigs
}

func (s *StaticServerList) GetServers() []loadbalancer.Server {
	servers := make([]loadbalancer.Server, len(s.Servers))

	for i, config := range s.Servers {
		host, portStr, err := net.SplitHostPort(config)

		port, err := strconv.Atoi(portStr)

		print(err) //TODO: deal with err

		servers[i] = loadbalancer.Server{Host: host, Port: port}
	}

	return servers
}

func (s *StaticServerList) GetServerKey() string {
	key := fmt.Sprintf("nuvem.loadbalancer.%s.serverlist.static.servers", s.Namespace)
	fmt.Printf("key %+v\n", key)
	return key
}

func (s *StaticServerList) GetNamespace() string {
	return s.Namespace
}

func NewStaticServerList() ServerList {
	return &StaticServerList{}
}

var StaticFactoryKey = "StaticServerList"

func init() {
	println("registering static serverlist")
	Register(StaticFactoryKey, NewStaticServerList)
}
