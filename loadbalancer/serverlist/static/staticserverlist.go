package static

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist/builder"
	"github.com/spf13/viper"
	"net"
	"strconv"
)

type (
	StaticServerList struct {
		namespace string
	}
)

func NewStaticServerList() serverlist.ServerList {
	return &StaticServerList{}
}

func (s *StaticServerList) Init(namespace string) {
	s.namespace = namespace
}

func (s *StaticServerList) GetServers() []loadbalancer.Server {
	key := fmt.Sprintf("loadbalancer.%s.serverlist.static.servers", s.namespace)
	fmt.Printf("key %+v\n", key)
	serverConfigs := viper.GetStringSlice(key)
	fmt.Printf("serverConfigs %+v\n", serverConfigs)

	servers := make([]loadbalancer.Server, len(serverConfigs))

	for i, config := range serverConfigs {
		host, portStr, err := net.SplitHostPort(config)

		port, err := strconv.Atoi(portStr)

		print(err) //TODO: deal with err

		servers[i] = loadbalancer.Server{Host: host, Port: port}
	}

	return servers
}

func init() {
	println("registering static lb")
	builder.Register("StaticServerList", NewStaticServerList)
}
