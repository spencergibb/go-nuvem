package registry

import "github.com/spencergibb/go-nuvem/util"

type Registry interface {
	util.Configurable
	Register(util.Instance)
	Unregister(util.Instance)
}

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Create(namespace string) Registry {
	result := factories.CallFactory("nuvem.registry", namespace, StaticFactoryKey)

	sl := result.Interface().(Registry)
	sl.Configure(namespace)

	return sl
}
