package registry

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/util"
	"github.com/spf13/viper"
)

func NewStaticBuilder(ns string) StaticBuilder {
	b := StaticBuilder{namespace: ns}
	return b
}

type StaticBuilder struct {
	namespace string
}

func (b StaticBuilder) Build() *StaticRegistry {
	registry := StaticRegistry{}
	registry.Namespace = b.namespace
	registry.Instances = make(map[string]util.Instance)
	return &registry
}

type StaticRegistry struct {
	Namespace string
	Instances map[string]util.Instance
}

func (s *StaticRegistry) Configure(namespace string) {
	if s.Namespace != "" {
		//TODO: use logging
		fmt.Errorf("%s already inited: %s", StaticFactoryKey, s.Namespace)
		return
	}
	s.Namespace = namespace
	s.Instances = make(map[string]util.Instance)
}

func (s *StaticRegistry) Register(instance util.Instance) {
	//TODO: error checking/warnings
	s.Instances[instance.Id] = instance
	s.setServers()
}

func (s *StaticRegistry) Unregister(instance util.Instance) {
	//TODO: error checking/warnings
	if _, ok := s.Instances[instance.Id]; ok {
		delete(s.Instances, instance.Id)
		s.setServers()
	}
}

func (s *StaticRegistry) GetNamespace() string {
	return s.Namespace
}

func (s *StaticRegistry) setServers() {
	var servers []string

	for _, instance := range s.Instances {
		server := fmt.Sprintf("%s:%d", instance.Host, instance.Port)
		servers = append(servers, server)
	}

	fmt.Printf("Setting static servers to %+v\n", servers)
	viper.SetDefault(util.GetStaticRegistryKey(s), servers)
}

func NewStaticRegistry() Registry {
	return &StaticRegistry{}
}

var StaticFactoryKey = "StaticRegistry"

func init() {
	println("registering static discovery")
	Register(StaticFactoryKey, NewStaticRegistry)
}
