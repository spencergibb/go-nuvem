package discovery

import (
	"github.com/spencergibb/go-nuvem/util"
)

type Discovery interface {
	util.Configurable
	GetIntances() []util.Instance
}

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Create(namespace string) Discovery {
	result := factories.CallFactory("nuvem.discovery", namespace, StaticFactoryKey)

	sl := result.Interface().(Discovery)
	sl.Configure(namespace)

	return sl
}
