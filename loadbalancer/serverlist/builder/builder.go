package builder

import (
	"github.com/spencergibb/go-nuvem/util"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
)

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Build(namespace string) serverlist.ServerList {
	result := factories.CallFactory("loadbalancer.serverlist", namespace, "StaticServerList")

	sl := result.Interface().(serverlist.ServerList)
	sl.Configure(namespace)

	return sl
}
