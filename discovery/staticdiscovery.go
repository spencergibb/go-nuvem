package discovery

import (
	"fmt"
	"github.com/spf13/viper"
	"net"
	"strconv"
)

type StaticDiscovery struct {
	Namespace string
	Instances []string
}

func (s *StaticDiscovery) Configure(namespace string) {
	if s.Namespace != "" {
		//TODO: use logging
		fmt.Errorf("%s already inited: %s", StaticFactoryKey, s.Namespace)
		return
	}
	s.Namespace = namespace
	instances := viper.GetStringSlice(s.GetServerKey())
	fmt.Printf("instances %+v\n", instances)
	s.Instances = instances
}

func (s *StaticDiscovery) GetIntances() []Instance {
	instances := make([]Instance, len(s.Instances))

	for i, config := range s.Instances {
		host, portStr, err := net.SplitHostPort(config)

		port, err := strconv.Atoi(portStr)

		print(err) //TODO: deal with err

		instances[i] = Instance{Host: host, Port: port}
	}

	return instances
}

func (s *StaticDiscovery) GetServerKey() string {
	key := fmt.Sprintf("nuvem.discovery.%s.static.servers", s.Namespace)
	fmt.Printf("key %+v\n", key)
	return key
}

func NewStaticDiscovery() Discovery {
	return &StaticDiscovery{}
}

var StaticFactoryKey = "StaticDiscovery"

func init() {
	println("registering static discovery")
	Register(StaticFactoryKey, NewStaticDiscovery)
}
