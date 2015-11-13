package builder

import (
	"github.com/spencergibb/go-nuvem/util"
//	"fmt"
//	"errors"
//	"strconv"
//	"reflect"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spf13/viper"
	"fmt"
//	"reflect"
)

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Build(namespace string) loadbalancer.LoadBalancer {
	key := fmt.Sprintf("loadbalancer.%s.factory", namespace)
	var factory string
	if !viper.IsSet(key) {
		//TODO: warn
		println("key is not set: ", key)
		factory = "NoopLoadBalancer" //TODO: default
	} else {
		factory = viper.GetString(key)
	}
	println("factory ", factory)
	results, err := factories.Call(factory)

	if len(results) != 1 {
		return nil//, errors.New("Wrong number of loadbalancer results " + strconv.Itoa(len(results)))
	}

	result := results[0]
//	fmt.Printf("Kind %+v\n", result.Kind())
//	fmt.Printf("isPtr %+v\n", result.Kind() == reflect.Ptr)
//	fmt.Printf("here %+v\n", result.Interface())

	lb := result.Interface().(loadbalancer.LoadBalancer)
//	switch t := lb.(type) {
//	default:
//		fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
//	case loadbalancer.LoadBalancer:
//		return lb.(loadbalancer.LoadBalancer), nil
//	}

	//.Interface().(LoadBalancer)
	lb.Init(namespace)

	print(err) //TODO deal with err
	return lb
	//	return lb, err
}
