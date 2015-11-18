package rule
import (
	"github.com/spencergibb/go-nuvem/util"
	"github.com/spencergibb/go-nuvem/loadbalancer"
)

type Rule interface {
	util.Configurable
	Choose(servers []loadbalancer.Server) *loadbalancer.Server
}