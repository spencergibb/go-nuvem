package factory

import (
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
	"github.com/spencergibb/go-nuvem/util"
)

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Create(namespace string) serverlist.ServerList {
	result := factories.CallFactory("loadbalancer.serverlist", namespace, "StaticServerList")

	sl := result.Interface().(serverlist.ServerList)
	sl.Configure(namespace)

	return sl
}
