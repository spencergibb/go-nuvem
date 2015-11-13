package loadbalancer

type (
	LoadBalancer interface {

		Choose() Server
	}

	// see http://www.captaincodeman.com/2015/03/05/dependency-injection-in-go-golang/
	// for explanation of DI in go
	loadBalancerFactory func(namespace string) LoadBalancer
)

var (
	New loadBalancerFactory
)