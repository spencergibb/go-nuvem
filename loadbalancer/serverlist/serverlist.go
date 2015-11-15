package serverlist

import (
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/util"
)

type ServerList interface {
	util.Initable
	GetServers() []loadbalancer.Server
}

func init() {
}
