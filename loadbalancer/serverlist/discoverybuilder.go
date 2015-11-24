package serverlist

import "github.com/spencergibb/go-nuvem/discovery"

func NewDiscoveryBuilder(ns string) DiscoveryBuilder {
	b := DiscoveryBuilder{namespace: ns}
	return b
}

type DiscoveryBuilder struct {
	namespace string
	discovery discovery.Discovery
}

func (b DiscoveryBuilder) Discovery(discovery discovery.Discovery) DiscoveryBuilder {
	b.discovery = discovery
	return b
}

func (b DiscoveryBuilder) Build() ServerList {
	serverList := DiscoveryServerList{}
	serverList.Namespace = b.namespace
	serverList.Discovery = b.discovery
	return &serverList
}
