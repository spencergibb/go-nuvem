package loadbalancer

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFactory(t *testing.T) {
	viper.SetConfigType("yaml")
	yaml := []byte(``)
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	viper.SetDefault("loadbalancer.test.factory", "NoopLoadBalancer")

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	lb := Create("test")
	assertLoadBalancer(t, lb)
}

func TestManual(t *testing.T) {
	println("\nTestBuilder")
	lb := NoopLoadBalancer{Namespace: "test"}

	assertLoadBalancer(t, &lb)
}

func assertLoadBalancer(t *testing.T, lb LoadBalancer) {
	require.NotNil(t, lb, "lb was nil")

	server := lb.Choose()
	assert.Nil(t, server, "server was not nil")
}
