package rule

import (
	"github.com/spencergibb/go-nuvem/util"
)

var factories = util.NewFuncs()

func RegisterRule(name string, fn interface{}) (err error) {
	return factories.Bind(name, fn)
}

func CreateRule(namespace string) Rule {
	result := factories.CallFactory("loadbalancer.rule", namespace, "RandomRule")

	rule := result.Interface().(Rule)
	rule.Configure(namespace)

	return rule
}
