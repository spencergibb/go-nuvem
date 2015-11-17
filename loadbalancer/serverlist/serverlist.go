package serverlist

import (
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/util"
)

type ServerList interface {
	util.Configurable
	GetServers() []loadbalancer.Server
}

type Builder interface {
	Namespace(ns string) Builder
	Build() ServerList
}

func init() {
}
