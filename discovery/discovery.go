package discovery

import (
	"github.com/spencergibb/go-nuvem/util"
)

type Instance struct {
	Id   string
	Host string
	Port int
}

type Discovery interface {
	util.Configurable
	GetIntances() []Instance
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
