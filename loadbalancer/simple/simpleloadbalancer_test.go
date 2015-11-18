package simple

import (
	"bytes"
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/rule"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFactory(t *testing.T) {
	viper.SetConfigType("yaml")
	yaml := []byte(`
loadbalancer.test.serverlist.static.servers:
- localhost:8080
- 127.0.0.1:9080
`)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("loadbalancer.test.factory", FactoryKey)

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	lb := loadbalancer.Create("test")
	assertLoadBalancer(t, lb)
}

func TestBuilder(t *testing.T) {
	println("\nTestBuilder")
	lb := NewBuilder().
		Namespace("test").
		ServerList(serverlist.NewStaticBuilder("test").Servers("10.0.0.1:80").Build()).
		Rule(rule.NewRandomRule()).
		Build()

	assertLoadBalancer(t, lb)
}

func assertLoadBalancer(t *testing.T, lb loadbalancer.LoadBalancer) {
	require.NotNil(t, lb, "lb was nil")

	server := lb.Choose()
	assert.NotNil(t, server, "server was nil")
}
