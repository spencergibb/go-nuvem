package serverlist

import (
	"github.com/spencergibb/go-nuvem/util"
)

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Create(namespace string) ServerList {
	result := factories.CallFactory("loadbalancer.serverlist", namespace, "StaticServerList")

	sl := result.Interface().(ServerList)
	sl.Configure(namespace)

	return sl
}
