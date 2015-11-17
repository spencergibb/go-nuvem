package static

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist/factory"
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
		fmt.Errorf("%s already inited: %s", FactoryKey, s.Namespace)
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
	key := fmt.Sprintf("loadbalancer.%s.serverlist.static.servers", s.Namespace)
	fmt.Printf("key %+v\n", key)
	return key
}

func NewStaticServerList() serverlist.ServerList {
	return &StaticServerList{}
}

var FactoryKey = "StaticServerList"

func init() {
	println("registering static serverlist")
	factory.Register(FactoryKey, NewStaticServerList)
}
