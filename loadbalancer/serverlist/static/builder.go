package static

import (
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
)

func NewBuilder(ns string) Builder {
	b := Builder{namespace: ns}
	return b
}

type Builder struct {
	namespace string
	servers   []string
}

func (b Builder) Servers(servers ...string) Builder {
	b.servers = servers
	return b
}

func (b Builder) Build() serverlist.ServerList {
	serverList := StaticServerList{}
	serverList.Namespace = b.namespace
	serverList.Servers = b.servers
	return &serverList
}
