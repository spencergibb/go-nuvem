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
	Servers   []string
}

func (b Builder) Build() serverlist.ServerList {
	serverList := StaticServerList{}
	serverList.Namespace = b.namespace
	serverList.Servers = b.Servers
	//	serverList.Init(b.namespace)
	return &serverList
}
