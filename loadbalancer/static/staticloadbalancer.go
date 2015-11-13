package static

import (
	"fmt"
    "github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/builder"
    "github.com/spf13/viper"
	"math/rand"
	"net"
	"strconv"
)

type (
	StaticLoadBalancer struct {
		namespace string
	}
)

func NewStaticLoadBalancer() loadbalancer.LoadBalancer {
	return &StaticLoadBalancer{}
}

func (s *StaticLoadBalancer) Init(namespace string) {
	s.namespace = namespace
}

func (s *StaticLoadBalancer) Choose() *loadbalancer.Server {
	key := fmt.Sprintf("loadbalancer.%s.static.servers", s.namespace)
//	fmt.Printf("key %+v\n", key)
	servers := viper.GetStringSlice(key)


	if (len(servers) == 0) {
		return nil
	}

	fmt.Printf("servers %+v\n", servers)
//	TODO: implement rules
	idx := rand.Intn(len(servers))

	host, portStr, err := net.SplitHostPort(servers[idx])

	port, err := strconv.Atoi(portStr)

	print(err) //TODO: deal with err

	return &loadbalancer.Server{Host: host, Port: port}
}

func init() {
	println("registering static lb")
	builder.Register("StaticLoadBalancer", NewStaticLoadBalancer)
}
