package rule

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"math/rand"
)

type RandomRule struct {
	Namespace string
}

func (s *RandomRule) Configure(namespace string) {
	if s.Namespace != "" {
		//TODO: use logging
		fmt.Errorf("%s already inited: %s", FactoryKey, s.Namespace)
		return
	}
	s.Namespace = namespace
}

func (s *RandomRule) Choose(servers []loadbalancer.Server) *loadbalancer.Server {
	idx := rand.Intn(len(servers))
	return &servers[idx]
}

var FactoryKey = "RandomRule"

func NewRandomRule() Rule {
	return &RandomRule{}
}

func init() {
	Register(FactoryKey, NewRandomRule)
}
