package discovery

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/util"
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
	instances := viper.GetStringSlice(util.GetStaticRegistryKey(s))
	fmt.Printf("instances %+v\n", instances)
	s.Instances = instances
}

func (s *StaticDiscovery) GetIntances() []util.Instance {
	instances := make([]util.Instance, len(s.Instances))

	for i, config := range s.Instances {
		host, portStr, err := net.SplitHostPort(config)

		port, err := strconv.Atoi(portStr)

		print(err) //TODO: deal with err

		instances[i] = util.Instance{Host: host, Port: port}
	}

	return instances
}

func (s *StaticDiscovery) GetNamespace() string {
	return s.Namespace
}

func NewStaticDiscovery() Discovery {
	return &StaticDiscovery{}
}

var StaticFactoryKey = "StaticDiscovery"

func init() {
	println("registering static discovery")
	Register(StaticFactoryKey, NewStaticDiscovery)
}
