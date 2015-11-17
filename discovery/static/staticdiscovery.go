package static

import (
	"github.com/spencergibb/go-nuvem/discovery"
)

type StaticDiscovery struct {
}

func (s StaticDiscovery) GetIntances() []discovery.Instance {
	return nil
}
