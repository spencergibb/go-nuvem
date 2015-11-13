package static

import (
    "github.com/spencergibb/go-nuvem/loadbalancer"
    "github.com/spf13/viper"
	"math/rand"
	"net"
	"strconv"
	"fmt"
)

type (
	staticLoadBalancer struct {
		namespace string
	}
)

func NewStaticLoadBalancer() loadbalancer.LoadBalancer {
	return &staticLoadBalancer{}
}

func (s *staticLoadBalancer) Init(namespace string) {
	s.namespace = namespace
}

func (s *staticLoadBalancer) Choose() loadbalancer.Server {
	servers := viper.GetStringSlice(fmt.Sprintf("loadbalancer.static.%s.servers", s.namespace))


	if (len(servers) == 0) {
		var s loadbalancer.Server
		return s
	}

//	fmt.Printf("%+v\n", servers)
//	TODO: implement rules
	idx := rand.Intn(len(servers))

	host, portStr, err := net.SplitHostPort(servers[idx])

	port, err := strconv.Atoi(portStr)

	print(err)

	return loadbalancer.Server{Host: host, Port: port}
}