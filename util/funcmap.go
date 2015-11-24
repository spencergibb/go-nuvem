package util

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

var (
	ErrParamsNotAdapted = errors.New("The number of params is not adapted.")
)

type Funcs map[string]reflect.Value

func NewFuncs() Funcs {
	return make(Funcs)
}

func (f Funcs) Bind(name string, fn interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(name + " is not callable.")
		}
	}()
	v := reflect.ValueOf(fn)
	v.Type().NumIn()
	f[name] = v
	return
}

func (f Funcs) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	if _, ok := f[name]; !ok {
		err = errors.New(name + " does not exist.")
		return
	}
	if len(params) != f[name].Type().NumIn() {
		err = ErrParamsNotAdapted
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f[name].Call(in)
	return
}

func (f Funcs) CallFactory(keyPrefix string, namespace string, defaultFactory string) *reflect.Value { //, error) {
	key := fmt.Sprintf("%s.%s.factory", keyPrefix, namespace)
	var factory string
	if !viper.IsSet(key) {
		//TODO: warn
		println("key is not set: ", key)
		factory = defaultFactory
	} else {
		println("key is set: ", key)
		factory = viper.GetString(key)
	}
	fmt.Printf("factory '%s'\n", factory)
	results, err := f.Call(factory)

	fmt.Errorf("error in CallFactory. key: %s, factory: %s\n%s", key, factory, err) //TODO deal with err

	if len(results) != 1 {
		return nil //, errors.New("Wrong number of loadbalancer results " + strconv.Itoa(len(results)))
	}

	result := results[0]
	return &result
}
