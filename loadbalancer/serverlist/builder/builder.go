package builder

import (
	"github.com/spencergibb/go-nuvem/util"
//	"fmt"
//	"errors"
//	"strconv"
//	"reflect"
	"github.com/spf13/viper"
	"fmt"
//	"reflect"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
)

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Build(namespace string) serverlist.ServerList {
	key := fmt.Sprintf("loadbalancer.serverlist.%s.factory", namespace)
	var factory string
	if !viper.IsSet(key) {
		//TODO: warn
		println("key is not set: ", key)
		factory = "StaticServerList" //TODO: default
	} else {
		factory = viper.GetString(key)
	}
	println("factory ", factory)
	results, err := factories.Call(factory)

	if len(results) != 1 {
		return nil//, errors.New("Wrong number of loadbalancer results " + strconv.Itoa(len(results)))
	}

	result := results[0]

	initable := result.Interface().(serverlist.ServerList)
	initable.Init(namespace)

	print(err) //TODO deal with err
	return initable
	//	return lb, err
}
