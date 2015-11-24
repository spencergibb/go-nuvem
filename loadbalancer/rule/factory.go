package rule

import (
	"github.com/spencergibb/go-nuvem/util"
)

var factories = util.NewFuncs()

func Register(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func Create(namespace string) Rule {
	result := factories.CallFactory("nuvem.loadbalancer.rule", namespace, "RandomRule")

	rule := result.Interface().(Rule)
	rule.Configure(namespace)

	return rule
}
