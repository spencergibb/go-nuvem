package simple

import (
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist/static"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

/*func TestFactory(t *testing.T) {
	viper.SetConfigType("yaml")
	yaml := []byte(``)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("loadbalancer.test.factory", "NoopLoadBalancer")

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	lb := factory.Create("test")
	assertLoadBalancer(t, lb)
}*/

func TestBuilder(t *testing.T) {
	println("\nTestBuilder")
	lb := NewBuilder().
		Namespace("test").
		ServerList(static.NewBuilder("test").Servers("10.0.0.1:80").Build()).
		Build()

	assertLoadBalancer(t, lb)
}

func assertLoadBalancer(t *testing.T, lb loadbalancer.LoadBalancer) {
	require.NotNil(t, lb, "lb was nil")

	server := lb.Choose()
	assert.NotNil(t, server, "server was nil")
}
