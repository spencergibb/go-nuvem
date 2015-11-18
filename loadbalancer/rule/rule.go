package rule

import (
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/util"
)

type Rule interface {
	util.Configurable
	Choose(servers []loadbalancer.Server) *loadbalancer.Server
}
